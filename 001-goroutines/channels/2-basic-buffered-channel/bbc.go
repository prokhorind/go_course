package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	defer close(ch)

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Дані відправлені у канал")

	// Отримуємо значення з каналу
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
