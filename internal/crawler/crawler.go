// Package crawler provides HTTP fetching, robots.txt parsing, and politeness
// enforcement for the distributed web crawler.
//
// This package implements the core crawling logic including:
//   - HTTP content fetching with configurable timeouts and retries
//   - robots.txt parsing and compliance
//   - Politeness policies (per-domain rate limiting)
//   - User-Agent management
//   - Redirect handling (301, 302, 307, 308)
//   - Content size limits
package crawler

// Fetcher handles HTTP requests to download web pages.
type Fetcher struct {
	UserAgent        string
	MaxContentSize   int64 // bytes
	RequestTimeout   int   // seconds
	MaxRedirects     int
	RetryCount       int
	RetryBackoffBase int // milliseconds
}

// NewFetcher creates a new Fetcher with default configuration.
func NewFetcher() *Fetcher {
	return &Fetcher{
		UserAgent:        "CrawlBot/1.0 (+contact@project.edu)",
		MaxContentSize:   10 * 1024 * 1024, // 10 MB
		RequestTimeout:   30,
		MaxRedirects:     5,
		RetryCount:       3,
		RetryBackoffBase: 1000,
	}
}

// RobotsChecker parses and evaluates robots.txt rules for a given domain.
type RobotsChecker struct {
	CacheTTLSeconds int
}

// NewRobotsChecker creates a new RobotsChecker with default TTL.
func NewRobotsChecker() *RobotsChecker {
	return &RobotsChecker{
		CacheTTLSeconds: 86400, // 24 hours
	}
}

// PolitenessEnforcer manages per-domain request rate limiting.
type PolitenessEnforcer struct {
	DefaultDelayMs     int
	MaxConcurrentPerDomain int
}

// NewPolitenessEnforcer creates a new PolitenessEnforcer with default settings.
func NewPolitenessEnforcer() *PolitenessEnforcer {
	return &PolitenessEnforcer{
		DefaultDelayMs:     1000,
		MaxConcurrentPerDomain: 2,
	}
}
