package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	
	// Define API endpoint - activate the PIR sensor
	const httpbin = "https://httpbin.org/post"

	// Make a POST request
	// creat the JSON request body:
	reqBody := strings.NewReader(`
		{
			"val": "this is a value",
			"num": 42
		}
	`)

	// post the generated request body to the camera API
	resp, err := http.Post(httpbin, "application/json", reqBody)
	if err != nil {
		return
	}

	// Read the response
	content, _ := ioutil.ReadAll(resp.Body)
	// Format the output
	fmt.Printf("%s\n", content)


	// Closing response
	defer resp.Body.Close()

	// Post form data
	data := url.Values{}
	data.Add("formField1", "Content of form field 1")
	data.Add("formField2", "Content of form field 2")
	respForm, err := http.PostForm(httpbin, data)
	contentForm, _ := ioutil.ReadAll(respForm.Body)
	defer respForm.Body.Close()
	fmt.Printf("%s\n", contentForm)
}