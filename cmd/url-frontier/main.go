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
	fmt.Printf("URL Frontier Service v%s (built %s)\n", version, buildTime)
	fmt.Println("Responsible for: priority queue management, domain-level politeness enforcement")
	fmt.Println("Status: scaffold — implementation pending")
	os.Exit(0)
}
