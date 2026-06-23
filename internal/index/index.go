// Package index implements the inverted index, forward index, tokenization,
// stemming, compression, and shard management for the search engine.
//
// Architecture (per SRS Section 8):
//   - Tokenizer: Unicode-aware word boundary tokenization
//   - Stemmer: Snowball (Porter2) algorithm
//   - InvertedIndex: Posting lists with delta encoding + VByte compression
//   - ForwardIndex: Document metadata and snippets
//   - ShardManager: Hash-based shard allocation and routing
package index

// InvertedIndex maps terms to compressed posting lists.
type InvertedIndex struct {
	ShardID      int
	NumDocuments uint32
}

// NewInvertedIndex creates a new inverted index for a shard.
func NewInvertedIndex(shardID int) *InvertedIndex {
	return &InvertedIndex{
		ShardID:      shardID,
		NumDocuments: 0,
	}
}

// Posting represents a single document occurrence of a term.
type Posting struct {
	DocID     uint32
	TermFreq  uint16
	Positions []uint32
}

// PostingList is a sorted list of postings for a term.
type PostingList struct {
	Term     string
	DocFreq  uint32
	Postings []Posting
}

// Tokenizer splits text into tokens using Unicode-aware word boundaries.
type Tokenizer struct{}

// NewTokenizer creates a new tokenizer.
func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

// Stemmer applies Snowball (Porter2) stemming algorithm.
type Stemmer struct {
	Language string
}

// NewStemmer creates a new stemmer for the specified language.
func NewStemmer(language string) *Stemmer {
	return &Stemmer{Language: language}
}

// ShardManager handles shard allocation and routing.
type ShardManager struct {
	NumShards     int
	ReplicaFactor int
}

// NewShardManager creates a new shard manager.
func NewShardManager(numShards, replicaFactor int) *ShardManager {
	return &ShardManager{
		NumShards:     numShards,
		ReplicaFactor: replicaFactor,
	}
}
