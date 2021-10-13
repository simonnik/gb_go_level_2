package main

import (
	"fmt"
	"sync"
)

const threads = 1000

func Worker() {
	var (
		counter int
		wg      = sync.WaitGroup{}
	)

	wg.Add(threads)
	for i := 0; i < threads; i += 1 {
		go func(i int) {
			defer wg.Done()

			counter += 1

			fmt.Println("Goroutine", i+1)
		}(i)
	}
	wg.Wait()

	fmt.Println("Goroutines completed:", counter)
}
