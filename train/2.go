package main

import (
	"context"
	"fmt"
	"time"
)

// На диске лежит очень большой файл объёмом более 100 Gb.
//  Файл содержит целые числа — по одному в каждой строке.
//  Нужно отсортировать их по возрастанию.
//  Опишите алгоритм решения такой задачи

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch := make(chan struct{})

	go func() {
		// делаем что-то
		time.Sleep(10 * time.Second)
		// если контекст отменен, вернем ошибку
		if ctx.Err() != nil {
			fmt.Println("Функция завершила работу")
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
