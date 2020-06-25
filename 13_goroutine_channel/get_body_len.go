package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{
		URL:  url,
		Size: len(body),
	}
}

func main() {
	sizeChannel := make(chan Page)

	urls := []string{"https://example.com/", "https://golang.org/doc", "https://google.com"}

	for _, url := range urls {
		go responseSize(url, sizeChannel)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizeChannel)
	}

}
