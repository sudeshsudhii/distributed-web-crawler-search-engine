// Package search implements search ranking algorithms including BM25, PageRank,
// HITS, vector search, hybrid retrieval (RRF), and cross-encoder reranking.
//
// Architecture (per SRS Section 9):
//   - BM25Scorer: Okapi BM25 with k1=1.2, b=0.75
//   - PageRankScorer: Power iteration over web graph
//   - HITSScorer: Hub/Authority computation on query subgraphs
//   - HybridSearcher: RRF fusion of BM25 + Vector results
//   - QueryParser: Boolean expression parsing (AND, OR, NOT)
package search

// BM25Scorer implements Okapi BM25 scoring.
type BM25Scorer struct {
	K1    float64 // Term frequency saturation (default: 1.2)
	B     float64 // Document length normalization (default: 0.75)
	AvgDL float64 // Average document length in corpus
}

// NewBM25Scorer creates a new BM25 scorer with default parameters.
func NewBM25Scorer(avgDL float64) *BM25Scorer {
	return &BM25Scorer{
		K1:    1.2,
		B:     0.75,
		AvgDL: avgDL,
	}
}

// SearchResult represents a single search result with multi-signal scores.
type SearchResult struct {
	DocID          uint32
	URL            string
	Title          string
	Snippet        string
	BM25Score      float64
	PageRankScore  float64
	VectorScore    float64
	HybridScore    float64
	FinalScore     float64
}

// PageRankScorer computes PageRank scores using power iteration.
type PageRankScorer struct {
	DampingFactor float64 // Default: 0.85
	MaxIterations int     // Default: 100
	Epsilon       float64 // Convergence threshold (default: 1e-6)
}

// NewPageRankScorer creates a new PageRank scorer with default parameters.
func NewPageRankScorer() *PageRankScorer {
	return &PageRankScorer{
		DampingFactor: 0.85,
		MaxIterations: 100,
		Epsilon:       1e-6,
	}
}

// HITSScorer computes hub and authority scores on query-specific subgraphs.
type HITSScorer struct {
	MaxIterations int
	Epsilon       float64
}

// NewHITSScorer creates a new HITS scorer.
func NewHITSScorer() *HITSScorer {
	return &HITSScorer{
		MaxIterations: 50,
		Epsilon:       1e-6,
	}
}

// HybridSearcher merges results from BM25 and Vector search using RRF.
type HybridSearcher struct {
	RRFConstant int // Default: 60
}

// NewHybridSearcher creates a new hybrid searcher.
func NewHybridSearcher() *HybridSearcher {
	return &HybridSearcher{
		RRFConstant: 60,
	}
}

// QueryParser parses search queries into structured query plans.
type QueryParser struct{}

// NewQueryParser creates a new query parser.
func NewQueryParser() *QueryParser {
	return &QueryParser{}
}
