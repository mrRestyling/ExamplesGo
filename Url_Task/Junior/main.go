package main

import "sync"

// Есть n-количество ссылок. Наша задача максимально быстро пройтись по всем ссылкам и получить http Code .
// Записать в мапу 200 или 500

// Важный поинт - паника при больше 1 млн операций (горутин + мютекс)
// Максимальное количество горутин конкурентно, которое можно запустить - 

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

	// 1.

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			u := sendRequest(url)
			mu.Lock()
			codes[u]++
			mu.Unlock()

		}()
	}

	wg.Wait()

}
