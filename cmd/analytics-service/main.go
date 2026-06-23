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
	fmt.Printf("Analytics Service v%s (built %s)\n", version, buildTime)
	fmt.Println("Responsible for: query logs, click-through analysis, crawl statistics")
	fmt.Println("Status: scaffold — implementation pending")
	os.Exit(0)
}
