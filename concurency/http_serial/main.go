package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN!\n", url)
	} else {
		defer res.Body.Close()
		fmt.Printf("%s -> status Code: %d \n", url, res.StatusCode)

		if res.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			file := strings.Split(url, "//")[1] //http://www.google.com
			file += ".txt"

			fmt.Printf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	wg.Done()
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}
	var wg sync.WaitGroup

	wg.Add(len(urls))
	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))
	}
	fmt.Println("No. of goroutine: ", runtime.NumGoroutine())
	wg.Wait()
}
