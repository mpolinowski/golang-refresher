package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func handleErr(err error) {
	if err !=nil {
		panic(err)
	}
}

// Check if file is already present
func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err !=nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Write data to file with `os`
func writeFileData() {
	// Create new file
	f, err := os.Create("os_create.txt")
	handleErr(err)
	// Close file afterwards
	defer f.Close()
	// Print name of generated file
	fmt.Println("File generated:", f.Name())
	// Write string to file
	f.WriteString("This is some text\n")
	// Write bytes to file
	//data2 := []byte("This some more text\n")
	data2 := []byte{'v', 'w', 'x', 'y', 'z', '\n'}
	f.Write(data2)
	// Sync everything to disk
	f.Sync()
}

// Append Data to file
func appendFileData(fname string, data string) {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	handleErr(err)
	defer f.Close()

	_, err = f.WriteString(data)
}

func main() {

		// Write data to file with `ioutil`
		// But skip if file is already present
		if !checkFileExists("datafile.txt") {
			data1 := []byte("Hi!\n")
			ioutil.WriteFile("datafile.txt", data1, 0644)
		}

		// Append data to file
		appendFileData("datafile.txt", "APPENDIX!")

		// Write data to file with `os`
		// But if file already exists
		if !checkFileExists("os_create.txt") {
			writeFileData()
		} else {
			// Trim file data to 22 bytes
			os.Truncate("os_create.txt", 22)
			fmt.Println("Log file has been truncated")
		}
	
}