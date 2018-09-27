package app

import (
	"fmt"
	"net/http"
	"time"
)

// Checker check site status
func Checker(sites []string) {
	c := make(chan string)

	for _, link := range sites {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is down")
		c <- link
		return
	}
	fmt.Println(link, " is up")
	c <- link
}
