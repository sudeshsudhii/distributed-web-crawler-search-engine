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
	fmt.Printf("Search API v%s (built %s)\n", version, buildTime)
	fmt.Println("Responsible for: query parsing, multi-shard scatter-gather, result aggregation")
	fmt.Println("Status: scaffold — implementation pending")
	os.Exit(0)
}
