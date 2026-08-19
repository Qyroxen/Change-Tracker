package main

import (
	"fmt"
	"os"
)

// change_tracker - Track system changes
func change_tracker(path string) {
	fmt.Println("========================================")
	fmt.Println("  Change-Tracker")
	fmt.Println("  Track system changes")
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
	change_tracker(path)
}
