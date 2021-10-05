package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Handle error
func handleErr(err error) {
	if err !=nil {
		panic(err)
	}
}

func main() {

	// Get default temp dir
	tempPath := os.TempDir()
	fmt.Println("Default temp dir:", tempPath)

	// Create a temporary file
	
	// Create temp data
	tempData := []byte("Some temp data")
	
	// Generate file with random filename
	tmpFile, err := ioutil.TempFile(tempPath, "temp_*.txt")
	if err != nil {
		panic(err)
	}
	
	// Write temp data to file
	if _, err := tmpFile.Write(tempData); err !=nil {
		panic(err)
	}
	
	// Read and print tempFile content
	data, _ := ioutil.ReadFile(tmpFile.Name())
	fmt.Printf("%s\n", data)

	// Remove temp file

	// Check if it is there
	fmt.Println("Remove tempfile:", tmpFile.Name())
	defer os.Remove(tmpFile.Name())


	// Create a randomly named temp dir inside the default tmp directory
	tmpDir, err := ioutil.TempDir("", "temp_*")
	if err != nil {
		panic(err)
	}
	// Remove temp dir
	fmt.Println("Remove temp dir:", tmpDir)
	defer os.RemoveAll(tmpDir)

}