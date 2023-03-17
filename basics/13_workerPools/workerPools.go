package main

import (
	"log"
	"net/http"
)

func main() {
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w := 1; w <= 3; w++ {
		go crawl(w, jobs, results)
	}
	var urls = []string{
		"https://google.com",
		"https://twitter.com",
		"https://github.com",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)
	for a := 1; a <= 4; a++ {
		result := <-results
		log.Println(result)
	}
}

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

func crawl(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("Worker ID: %d\n", wId)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{
			URL:    site.URL,
			Status: resp.StatusCode,
		}
	}
}
