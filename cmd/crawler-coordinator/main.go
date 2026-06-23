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
	fmt.Printf("Crawler Coordinator v%s (built %s)\n", version, buildTime)
	fmt.Println("Responsible for: crawl job management, worker orchestration, task distribution")
	fmt.Println("Status: scaffold — implementation pending")
	os.Exit(0)
}
