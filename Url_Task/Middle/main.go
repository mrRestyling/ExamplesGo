package main

import "sync"

// Есть n-количество ссылок. Наша задача максимально быстро пройтись по всем ссылкам и получить http Code .
// Записать в мапу 200 или 500

// Важный поинт - паника при больше 1 млн операций (горутин + мютекс)

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
