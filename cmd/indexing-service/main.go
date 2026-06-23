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
	fmt.Printf("Indexing Service v%s (built %s)\n", version, buildTime)
	fmt.Println("Responsible for: tokenization, stemming, inverted index construction, embedding generation")
	fmt.Println("Status: scaffold — implementation pending")
	os.Exit(0)
}
