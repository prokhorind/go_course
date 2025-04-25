package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int = 0
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				counter++
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for j := 0; j < 500; j++ {
				counter--
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Результат лічильника:", counter)
}
