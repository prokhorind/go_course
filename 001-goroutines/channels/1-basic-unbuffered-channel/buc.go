package main

import "fmt"

func main() {
	ch := make(chan string)
	defer close(ch)

	go func() {
		ch <- "Привіт від горутини!"
	}()

	msg := <-ch
	fmt.Println(msg)
}
