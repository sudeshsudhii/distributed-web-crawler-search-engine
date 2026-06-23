package frontier

import (
	"math"
	"testing"
)

func TestBloomFilterMath(t *testing.T) {
	// Testing 10M items with 1% FPR
	expectedItems := uint64(10000000)
	fpr := 0.01

	m := -float64(expectedItems) * math.Log(fpr) / math.Pow(math.Log(2), 2)
	numBits := uint64(math.Ceil(m))

	k := float64(numBits) / float64(expectedItems) * math.Log(2)
	numHashes := uint(math.Ceil(k))

	// Verify against known expected sizes for 1% FPR
	// m should be roughly 95.85 million bits (~11.4 MB)
	// k should be 7
	if numHashes != 7 {
		t.Errorf("expected 7 hash functions, got %d", numHashes)
	}

	expectedMB := float64(numBits) / 8 / 1024 / 1024
	if expectedMB < 11.0 || expectedMB > 12.0 {
		t.Errorf("expected ~11.4MB memory usage, got %.2f MB", expectedMB)
	}
}

func TestDomainRouter(t *testing.T) {
	router := NewDomainRouter(256)

	q1 := router.GetQueueName("example.com")
	q2 := router.GetQueueName("example.com")

	if q1 != q2 {
		t.Errorf("consistent hashing failed: %s != %s", q1, q2)
	}
	
	q3 := router.GetQueueName("github.com")
	if q1 == q3 {
		t.Logf("different domains happened to hash to the same queue %s, which is valid but unlikely with high queue counts", q1)
	}
}
