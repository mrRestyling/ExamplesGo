package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	Status int
}

func Crawl(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("Worker ID: %d\n", wId)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{Status: resp.StatusCode}
	}
}

func main() {
	fmt.Println("Worker pools")

	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w := 1; w <= 3; w++ {
		go Crawl(w, jobs, results)
	}

	urls := []string{
		"https://github.com/mrRestyling",
		"https://yandex.ru/",
		"https://yandex.ru/",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for a := 1; a <= 3; a++ {
		result := <-results
		log.Println(result)
	}

}
