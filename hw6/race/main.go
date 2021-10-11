// Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”
package main

import (
	"fmt"
	"sync"
)

const threads = 1000

func main() {
	var (
		count int
		wg    sync.WaitGroup
	)
	wg.Add(threads)
	for i := 0; i < threads; i += 1 {
		go func() {
			defer wg.Done()
			count += 1
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
