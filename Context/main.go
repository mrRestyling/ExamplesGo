package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch := make(chan struct{})

	go func() {
		// делаем что-то
		time.Sleep(2 * time.Second)
		// если контекст отменен, вернем ошибку
		if ctx.Err() != nil {
			fmt.Println("Функция завершила работу с ошибкой")
			return
		}
		// если контекст не отменен, продолжаем работать
		fmt.Println("Функция завершила работу")
		close(ch)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Контекст отменен")
	case <-ch:
		fmt.Println("Горутина завершила работу")
	}
}
