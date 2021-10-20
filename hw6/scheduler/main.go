// Написать многопоточную программу, в которой будет использоваться явный вызов планировщика.
// Выполните трассировку программы
package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	err := trace.Start(os.Stderr)

	if err != nil {
		log.Fatalf("Error %v", err)
	}
	var wg sync.WaitGroup

	defer trace.Stop()

	runtime.GOMAXPROCS(4)
	runtime.Gosched()

	go func() {
		defer wg.Done()
		wg.Add(3)
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			wg.Add(1)
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
