// Package frontier implements the distributed URL frontier with priority queuing,
// Bloom filter deduplication, and domain-aware politeness enforcement.
//
// Architecture (per SRS Section 6.1):
//   - Front Queue: Priority-based URL scheduling (Redis Sorted Sets)
//   - Router: Consistent hash routing by domain
//   - Back Queue: Per-domain politeness queues
//   - Bloom Filter: Probabilistic URL deduplication (10M URLs, 1% FPR)
package frontier

import (
	"context"
	"fmt"
	"hash/crc32"
	"math"

	"github.com/redis/go-redis/v9"
	"github.com/spaolacci/murmur3"
)

// PriorityQueue manages URL prioritization using Redis Sorted Sets.
type PriorityQueue struct {
	client *redis.Client
	key    string
}

// NewPriorityQueue creates a new priority queue backed by Redis.
func NewPriorityQueue(client *redis.Client, queueName string) *PriorityQueue {
	return &PriorityQueue{
		client: client,
		key:    queueName,
	}
}

// Push adds a URL to the priority queue with the given priority.
// Higher priority means the URL should be crawled sooner. In Redis sorted sets,
// we'll retrieve elements with the highest score.
func (pq *PriorityQueue) Push(ctx context.Context, url string, priority float64) error {
	return pq.client.ZAdd(ctx, pq.key, redis.Z{
		Score:  priority,
		Member: url,
	}).Err()
}

// Pop removes and returns the URL with the highest priority.
// Returns an empty string and redis.Nil error if the queue is empty.
func (pq *PriorityQueue) Pop(ctx context.Context) (string, error) {
	res, err := pq.client.ZPopMax(ctx, pq.key, 1).Result()
	if err != nil {
		return "", err
	}
	if len(res) == 0 {
		return "", redis.Nil
	}
	return res[0].Member.(string), nil
}

// BloomFilter provides probabilistic membership testing for visited URLs.
// Uses pipelined Redis SETBIT/GETBIT operations to reduce network roundtrips.
type BloomFilter struct {
	client    *redis.Client
	key       string
	numHashes uint
	numBits   uint64
}

// NewBloomFilter creates a new Redis-backed Bloom filter.
// It calculates the optimal number of hash functions (k) and bit array size (m).
func NewBloomFilter(client *redis.Client, key string, expectedItems uint64, fpr float64) *BloomFilter {
	// Calculate m = -n * ln(p) / (ln(2)^2)
	m := -float64(expectedItems) * math.Log(fpr) / math.Pow(math.Log(2), 2)
	numBits := uint64(math.Ceil(m))

	// Calculate k = (m/n) * ln(2)
	k := float64(numBits) / float64(expectedItems) * math.Log(2)
	numHashes := uint(math.Ceil(k))

	return &BloomFilter{
		client:    client,
		key:       key,
		numHashes: numHashes,
		numBits:   numBits,
	}
}

// Add adds a URL to the Bloom filter using a Redis pipeline.
func (bf *BloomFilter) Add(ctx context.Context, url string) error {
	pipe := bf.client.Pipeline()

	h1, h2 := murmur3.Sum128([]byte(url))

	for i := uint(0); i < bf.numHashes; i++ {
		// Kirsch-Mitzenmacher optimization: hash_i = h1 + i * h2
		combinedHash := h1 + uint64(i)*h2
		bitOffset := int64(combinedHash % bf.numBits)

		pipe.SetBit(ctx, bf.key, bitOffset, 1)
	}

	_, err := pipe.Exec(ctx)
	return err
}

// Exists checks if a URL might be in the Bloom filter.
func (bf *BloomFilter) Exists(ctx context.Context, url string) (bool, error) {
	pipe := bf.client.Pipeline()

	h1, h2 := murmur3.Sum128([]byte(url))

	var cmds []*redis.IntCmd
	for i := uint(0); i < bf.numHashes; i++ {
		combinedHash := h1 + uint64(i)*h2
		bitOffset := int64(combinedHash % bf.numBits)

		cmd := pipe.GetBit(ctx, bf.key, bitOffset)
		cmds = append(cmds, cmd)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	// If all bits are 1, it exists
	for _, cmd := range cmds {
		val, err := cmd.Result()
		if err != nil {
			return false, err
		}
		if val == 0 {
			return false, nil
		}
	}

	return true, nil
}

// DomainRouter assigns URLs to back queues using consistent hashing on domain.
type DomainRouter struct {
	NumQueues int
}

// NewDomainRouter creates a new domain router with specified number of queues.
func NewDomainRouter(numQueues int) *DomainRouter {
	return &DomainRouter{
		NumQueues: numQueues,
	}
}

// GetQueueName assigns a domain to one of the numbered back queues using CRC32.
func (dr *DomainRouter) GetQueueName(domain string) string {
	hash := crc32.ChecksumIEEE([]byte(domain))
	queueIndex := hash % uint32(dr.NumQueues)
	return fmt.Sprintf("queue:back:%d", queueIndex)
}
