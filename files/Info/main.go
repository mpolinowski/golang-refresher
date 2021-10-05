package main

import (
	"fmt"
	"os"
)

// Make sure a file is present
func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err !=nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {

	// Check file
	exists := checkFileExists("text.txt")
	fmt.Println("Is text.txt present:", exists)

	// Get file stats
	stats, err := os.Stat("text.txt")
	if err != nil {
		panic(err)
	}
	// Print modification time
	fmt.Println("Modification time:", stats.ModTime())
	// Print filesize
	fmt.Println("File size:", stats.Size(), "bytes")
	// Check file permissions
	fmt.Println("File Mode:", stats.Mode())
	// Check if it is a directory
	fmt.Println("Is it a directory:", stats.IsDir())
	// Check if it is a file
	fmode := stats.Mode()
	if fmode.IsRegular() {
		fmt.Println("This is a file not a symlink")
	}
}