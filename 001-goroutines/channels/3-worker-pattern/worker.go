package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Робітник %d обробляє завдання %d\n", id, job)
		time.Sleep(1 * time.Second) // Імітація роботи
		fmt.Printf("Робітник %d завершив завдання %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs) // Запускаємо горутини-робітники
	}

	for j := 1; j <= 5; j++ {
		jobs <- j // Відправляємо завдання у канал
	}
	close(jobs) // Закриваємо канал
}
