package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int = 0
	var wg sync.WaitGroup
	var mu sync.Mutex // М'ютекс

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for j := 0; j < 500; j++ {
				mu.Lock()
				counter--
				mu.Unlock()
			}
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("Результат лічильника:", counter)
}
