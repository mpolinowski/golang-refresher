package main

import (
	"fmt"
	"net/url"
)

func main() {
	
	// Define a url
	u := "https://mpolinowski.github.io:80/?chapter_filter=%22Dev+Notes%22&type_filter=%22Note%22&q=%22Golang%22&tag_filter=%5B%22Golang%22%5D"

	// Parse url into it's parts
	result, _ := url.Parse(u)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// Extract query's into variables
	val := result.Query()
	fmt.Println("Search:", val["q"])
	fmt.Println("Tags:", val["tag_filter"])
	fmt.Println("Types:", val["type_filter"])
	fmt.Println("Chapters:", val["chapter_filter"])

	// Create URL from components
	newURL := &url.URL {
		Scheme: "https",
		Host: "mpolinowski.github.io",
		Path: "/devnotes",
		RawQuery: "usr=admin&pwd=password",
	}

	// Print created url
	s := newURL.String()
	fmt.Println(s)
	// Modify url
	newURL.RawQuery = "usr=user&pwd=1234"
	s = newURL.String()
	fmt.Println(s)

	// Generate new url queries
	newvals := url.Values{}
	newvals.Add("pwd", "nopwd")
	newvals.Add("usr", "visitor")

	newURL.RawQuery = newvals.Encode()
	s = newURL.String()
	fmt.Println(s)
	
}