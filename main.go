package main

import (
	"fmt"
	"os"
)

// duplicate_finder - Find duplicate files
func duplicate_finder(path string) {
	fmt.Println("========================================")
	fmt.Println("  Duplicate-Finder")
	fmt.Println("  Find duplicate files")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	duplicate_finder(path)
}
