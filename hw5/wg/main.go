// Напишите программу, которая запускает n потоков и дожидается завершения их всех
package main

import (
	"fmt"
	"sync"
)

const threads = 100

func main() {
	var (
		count int
		wg    = sync.WaitGroup{}
	)

	wg.Add(threads)

	for i := 0; i < threads; i++ {
		go func() {
			count++
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Completed", count)
}
