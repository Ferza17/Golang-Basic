package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func checkURL(url string, c chan string) {
	res, err := http.Get(url)
	if err != nil {
		// fmt.Println(err)
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error : %v\n", err)
		c <- url // Sending to the channel
		fmt.Println(s)
	} else {
		defer res.Body.Close()
		s := fmt.Sprintf("%s -> status Code: %d \n", url, res.StatusCode)

		s += fmt.Sprintf("%s is UP\n", url)
		fmt.Println(s)
		c <- url
	}
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	// 1.
	c := make(chan string)
	for _, url := range urls {
		go checkURL(url, c)
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("No. of goroutines: ", runtime.NumGoroutine())

	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkUrl(<-c, c)
	// 	fmt.Println(strings.Repeat("#", 30))
	// 	time.Sleep(time.Second * 2)
	// }

	// for url := range c {
	// 	time.Sleep(time.Second * 2)
	// 	go checkURL(url, c)
	// }

	for url := range c {
		go func(u string) {
			time.Sleep(time.Second * 2)
			checkURL(u, c)
		}(url)
	}

}
