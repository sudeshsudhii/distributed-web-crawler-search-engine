// Package dedup implements content deduplication using SimHash and MinHash
// algorithms for near-duplicate detection in the crawl pipeline.
//
// Architecture (per SRS Section 7.4):
//   - SimHash: 64-bit fingerprints with Hamming distance ≤ 3 threshold
//   - MinHash: Jaccard similarity estimation with LSH for scalable comparison
//   - Two-tier strategy: Bloom filter (URL) + SimHash/MinHash (content)
package dedup

// SimHasher generates 64-bit SimHash fingerprints for text content.
type SimHasher struct {
	HammingThreshold int // Default: 3
}

// NewSimHasher creates a new SimHash fingerprint generator.
func NewSimHasher() *SimHasher {
	return &SimHasher{
		HammingThreshold: 3,
	}
}

// MinHasher generates MinHash signatures for Jaccard similarity estimation.
type MinHasher struct {
	NumHashFunctions int     // Default: 128
	SimilarityThreshold float64 // Default: 0.8
}

// NewMinHasher creates a new MinHash generator.
func NewMinHasher() *MinHasher {
	return &MinHasher{
		NumHashFunctions:    128,
		SimilarityThreshold: 0.8,
	}
}

// ContentFingerprint represents a document's deduplication fingerprint.
type ContentFingerprint struct {
	DocID         uint32
	SimHashValue  uint64
	MinHashSignature []uint64
}
