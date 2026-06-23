// Package frontier implements the distributed URL frontier with priority queuing,
// Bloom filter deduplication, and domain-aware politeness enforcement.
//
// Architecture (per SRS Section 6.1):
//   - Front Queue: Priority-based URL scheduling (Redis Sorted Sets)
//   - Router: Consistent hash routing by domain
//   - Back Queue: Per-domain politeness queues
//   - Bloom Filter: Probabilistic URL deduplication (10M URLs, 1% FPR)
package frontier

// PriorityQueue manages URL prioritization using Redis Sorted Sets.
type PriorityQueue struct {
	NumPriorityLevels int
}

// NewPriorityQueue creates a new priority queue.
func NewPriorityQueue(levels int) *PriorityQueue {
	return &PriorityQueue{
		NumPriorityLevels: levels,
	}
}

// BloomFilter provides probabilistic membership testing for visited URLs.
// Parameters: expected items = 10M, false positive rate = 1% → ~11.5 MB.
type BloomFilter struct {
	ExpectedItems     uint64
	FalsePositiveRate float64
}

// NewBloomFilter creates a new Bloom filter with specified parameters.
func NewBloomFilter(expectedItems uint64, fpr float64) *BloomFilter {
	return &BloomFilter{
		ExpectedItems:     expectedItems,
		FalsePositiveRate: fpr,
	}
}

// DomainRouter assigns URLs to back queues using consistent hashing on domain.
type DomainRouter struct {
	VirtualNodes int
}

// NewDomainRouter creates a new domain router with specified virtual nodes.
func NewDomainRouter(vnodes int) *DomainRouter {
	return &DomainRouter{
		VirtualNodes: vnodes,
	}
}
