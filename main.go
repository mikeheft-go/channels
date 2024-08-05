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

	c := make(chan string)

	for _, link := range links {
		go getStatus(c, link)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func getStatus(c chan string, link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("[ERROR]: ", err)
		c <- "[ERROR]: Might be down"
		os.Exit(1)
	}

	fmt.Println(link, "is up!")
	c <- "It's up!"
}
