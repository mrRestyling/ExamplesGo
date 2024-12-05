package main

import "sync"

// Добавляем WorkerPool

func sendRequest(url string) (code int) {
	// some code
	return 0
}

func main() {

	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	urls := make([]string, 100000000)
	codes := make(map[int]int)
	ch := make(chan string)

	go func() {
		for _, url := range urls {
			ch <- url
		}
	}()

	wg.Add(5)
	for range 5 {
		go func() {
			defer wg.Done()
			res := sendRequest(<-ch)

			mu.Lock()
			codes[res]++
			mu.Unlock()
		}()

	}
	wg.Wait()

}
