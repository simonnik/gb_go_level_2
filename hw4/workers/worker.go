package main

import (
	"fmt"
	"time"
)

func main() {
	limitWorkers := 1000
	total := 0
	var jobs = make(chan int, limitWorkers)

	for i := 0; i < limitWorkers; i++ {
		jobs <- 1
		go func() {
			total += <-jobs
		}()

	}

	time.Sleep(1 * time.Second)
	fmt.Println("total", total)
}
