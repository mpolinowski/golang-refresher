package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Handle error
func handleErr(err error) {
	if err !=nil {
		panic(err)
	}
}
// Get length of file
func getFileLength(filepath string) int64 {
	f, err := os.Stat(filepath)
	handleErr(err)
	return f.Size()
}

func main() {

	// Get length of file
	length := getFileLength("text.txt")
	fmt.Println("File length:", length)

	// Read entire file and print
	content, err := ioutil.ReadFile("text.txt")
	handleErr(err)
	fmt.Println("Entire Content:\n", string(content))
	fmt.Println()

	// Read file in chunks of 20 bytes
	// open the file
	const BuffSize = 20
	f, _ := os.Open("text.txt")
	defer f.Close()
	// Create buffer with size BuffSize
	b1 := make([]byte, BuffSize)
	fmt.Println("Content read in 20bytes chunks")
	for {
		n, err := f.Read(b1)
		// Continue looping
		if err != nil {
			// Until you reach the end of file
			if err != io.EOF {
				// If error other than EOF, panic
				handleErr(err)
			}
			// Break
			break
		}

		fmt.Println("Bytes read:", n)
		fmt.Println("Content:", string(b1[:n]))
	}

	// Read from a specific position
	// Create buffer to read 10 bytes
	b2 := make([]byte, 10)
	// Read the last 10 bytes from open file
	_, err = f.ReadAt(b2,length-int64(len(b2)))
	handleErr(err)
	fmt.Println()
	fmt.Println("Last 10bytes of file:", string(b2))

}