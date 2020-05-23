package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{

		"http://yahoo.com",
		"http://shib.me",
		"http://cloudflare.com",
		"http://youtube.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://google.com",
		"http://freshworks.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {

		go func(link string) {
			time.Sleep(time.Second * 10)
			checkLink(link, c)
		}(l)

	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, " is UP")
	c <- link
}
