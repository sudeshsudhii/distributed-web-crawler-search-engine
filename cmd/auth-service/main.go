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
	fmt.Printf("Auth Service v%s (built %s)\n", version, buildTime)
	fmt.Println("Responsible for: JWT issuance, validation, RBAC enforcement")
	fmt.Println("Status: scaffold — implementation pending")
	os.Exit(0)
}
