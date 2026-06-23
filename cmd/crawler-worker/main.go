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
	fmt.Printf("Crawler Worker v%s (built %s)\n", version, buildTime)
	fmt.Println("Responsible for: HTTP fetching, DNS resolution, raw content download")
	fmt.Println("Status: scaffold — implementation pending")
	os.Exit(0)
}
