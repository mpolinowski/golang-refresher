package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Create a directory
	os.Mkdir("newdir", os.ModePerm)

	// Create nested dir
	os.MkdirAll("newdir2/surprise", os.ModePerm)

	// Remove a directory
	defer os.Remove("newdir")

	// Remove directory and children
	defer os.RemoveAll("newdir2")

	// Read what the working directory is
	dir, _ := os.Getwd()
	fmt.Println("Current working dir:", dir)

	// Read the directory of the current process
	exedir, _ := os.Executable()
	fmt.Println("Process dir:", exedir)

	// Read and list content of current directory
	content, _ := ioutil.ReadDir(".")
	fmt.Println("Dir content:", content)
	for _, fi := range content {
		fmt.Println(fi.Name(), fi.IsDir())
	}
}