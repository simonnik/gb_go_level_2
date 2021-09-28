package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	handlePanic()
	createFiles()
	panicInGoroutine()
}

type ErrorWithTimestamp struct {
	message   string
	timestamp string
}

func NewErrorWithTimestamp(text string) error {
	return &ErrorWithTimestamp{
		message:   text,
		timestamp: time.Now().String(),
	}
}

func (e *ErrorWithTimestamp) Error() string {
	return fmt.Sprintf("error: %s\noccur at: %s \n", e.message, e.timestamp)
}

func handlePanic() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Handle panic:\n%v\n", NewErrorWithTimestamp(fmt.Sprintf("Error from panic %v", v)))
		}
	}()

	fmt.Println("### A panic example ###")

	var a int
	fmt.Println(5 / a)
}

func createFiles() {
	fmt.Println("### Creating 1M of empty files ###")
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovered from panic in createFiles:", v)
		}
	}()
	dir := "../tmpFiles"

	_ = os.Mkdir(dir, 0700)

	n := 1_000_000
	ind := 50000

	for i := 1; i <= n; i++ {
		f := createFile(fmt.Sprintf("%s/file_%d.txt", dir, i))
		func(f *os.File) {
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					log.Fatalf("error: %v\n", err)
				}
			}(f)
		}(f)

		if i%ind == 0 {
			fmt.Printf("Created %v\n", i)
		}
	}
	fmt.Println("Reading dir...")

	files, _ := ioutil.ReadDir(dir)
	fmt.Printf("Created %v files.\n", len(files))

	fmt.Println("Cleaning up...")
	err := os.RemoveAll(dir)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func createFile(pathFile string) *os.File {
	f, err := os.Create(pathFile)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		panic(err)
	}

	return f
}

func panicInGoroutine() {
	fmt.Println("### Handling panic in goroutine ###")

	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("Recovered from panic in goroutine:", v)
			}
		}()

		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)
}
