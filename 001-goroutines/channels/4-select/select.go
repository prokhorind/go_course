package main

import (
	"fmt"
	"time"
)

func worker(ch chan int, done chan bool) {
	for {
		select {
		case val := <-ch:
			fmt.Println("Отримано з каналу:", val)
		case <-done:
			fmt.Println("Завершення горутини")
			return
		default:
			fmt.Println("Немає доступних даних")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ch := make(chan int, 2)
	done := make(chan bool)

	go worker(ch, done)

	ch <- 42
	ch <- 43
	time.Sleep(1 * time.Second)
	done <- true // Завершення горутини
}
