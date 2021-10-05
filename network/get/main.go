package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	
	// Define API endpoint - get the PIR motion sensor state:
	const cameraAPI = "http://192.168.2.77:8090/param.cgi?cmd=getpirattr&usr=admin&pwd=instar"

	// Make GET request
	resp, err := http.Get(cameraAPI)
	if err != nil {
		return
	}

	// Closing response
	defer resp.Body.Close()

	// Splitting up the response
	fmt.Println("Status:", resp.Status)
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Protocol:", resp.Proto)
	fmt.Println("Content Length:", resp.ContentLength)

	// Build content from received bytes
	var sb strings.Builder
	content, _ := ioutil.ReadAll(resp.Body)
	bytecount, _ := sb.Write(content)
	// Format the output
	fmt.Println(bytecount, sb.String())

	
}