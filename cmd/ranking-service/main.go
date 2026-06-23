package main

import (
	"fmt"
	"os"
)

var (
	version   = "dev"
	buildTime = "unknown"
)

func main() {
	fmt.Printf("Ranking Service v%s (built %s)\n", version, buildTime)
	fmt.Println("Responsible for: BM25, PageRank, HITS, hybrid scoring, reranking")
	fmt.Println("Status: scaffold — implementation pending")
	os.Exit(0)
}
