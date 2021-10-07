// Реализуйте функцию для разблокировки мьютекса с помощью defer
package main

import (
	"fmt"
	"sync"
)

const threads = 100

func main() {
	var (
		count int
		m     = sync.Mutex{}
		wg    sync.WaitGroup
	)
	wg.Add(threads)

	for i := 0; i < threads; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			count++
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Completed", count)
}
