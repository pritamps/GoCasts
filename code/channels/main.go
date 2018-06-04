package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	// Communication will be of type string
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
