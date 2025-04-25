package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Привіт, горутина!")
}

func main() {
	go sayHello()

	go func() {
		fmt.Println("Анонімна горутина працює!")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Головна функція завершена")
}
