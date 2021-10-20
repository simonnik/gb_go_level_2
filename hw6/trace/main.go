//  Написать программу, которая использует мьютекс для безопасного доступа к данным
// из нескольких потоков. Выполните трассировку программы
package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
)

const threads = 100

func main() {
	err := trace.Start(os.Stderr)

	if err != nil {
		log.Fatalf("Error %v", err)
	}
	defer trace.Stop()

	var (
		count int
		m     = sync.Mutex{}
		wg    sync.WaitGroup
	)
	wg.Add(threads)

	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()
			m.Lock()
			defer m.Unlock()
			count++
		}()
	}
	wg.Wait()

	fmt.Println("Completed", count)
}
