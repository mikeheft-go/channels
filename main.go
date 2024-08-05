package main

import (
	"fmt"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		getStatus(link)
	}
}

func getStatus(link string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("[ERROR]: ", err)
		os.Exit(1)
	}

	fmt.Println(resp.StatusCode)
}
