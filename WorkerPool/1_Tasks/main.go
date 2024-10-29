package main

import (
	"fmt"
	"sync"
)

// Task представляет собой задачу, которую нужно выполнить
type Task struct {
	ID   int
	Data string
}

// Worker представляет собой рабочую горутину
type Worker struct {
	ID      int
	TaskCh  chan Task
	QuitCh  chan bool
	Results chan Task
}

// NewWorker создает новую рабочую горутину
func NewWorker(id int) *Worker {
	return &Worker{
		ID:      id,
		TaskCh:  make(chan Task),
		QuitCh:  make(chan bool),
		Results: make(chan Task),
	}
}

// Start запускает рабочую горутину
func (w *Worker) Start() {
	go func() {
		for {
			select {
			case task := <-w.TaskCh:
				// Выполняем задачу
				fmt.Printf("Рабочая горутина %d: выполняем задачу %d\n", w.ID, task.ID)
				// Отправляем результат
				w.Results <- task
			case <-w.QuitCh:
				// Останавливаем горутину
				return
			}
		}
	}()
}

// Stop останавливает рабочую горутину
func (w *Worker) Stop() {
	close(w.QuitCh)
}

func main() {
	// Создаем пул рабочих горутин
	workerPool := make([]*Worker, 4)
	for i := range workerPool {
		workerPool[i] = NewWorker(i + 1)
		workerPool[i].Start()
	}

	// Создаем задачи
	tasks := []Task{
		{ID: 1, Data: "Задача 1"},
		{ID: 2, Data: "Задача 2"},
		{ID: 3, Data: "Задача 3"},
		{ID: 4, Data: "Задача 4"},
	}

	// Отправляем задачи всем рабочим горутинам
	var wg sync.WaitGroup

	wg.Add(len(tasks))

	for _, task := range tasks {
		worker := workerPool[0]
		worker.TaskCh <- task
		workerPool = append(workerPool[1:], worker)

		go func() {
			defer wg.Done()
			result := <-worker.Results
			fmt.Printf("Результат задачи %d: %s\n", result.ID, result.Data)
		}()
	}

	// Ждем завершения всех рабочих горутин
	wg.Wait()

	// Останавливаем рабочие горутины
	for _, worker := range workerPool {
		worker.Stop()
	}
}
