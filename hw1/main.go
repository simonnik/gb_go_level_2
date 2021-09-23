package main

import (
	"fmt"
	"time"
)

func main() {
	handlePanic()
	//createFiles()
	panicInGoroutine()
}

type ErrorWithTimestamp struct {
	message   string
	timestamp string
}

func New(text string) error {
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
			err := New(fmt.Sprintf("Error from panic %v", v))
			fmt.Printf("Handle panic:\n%v\n", err)
		}
	}()

	fmt.Println("### A panic example ###")

	var a int
	fmt.Println(5 / a)
}

//func createFiles() {
//	fmt.Println("### Creating 1M of empty files ###")
//	dir := "../tmpFiles"
//
//	_ = os.Mkdir(dir, 0700)
//
//	n := 1_000_000
//
//	for i := 1; i <= n; i++ {
//		f.CreateAndCloseFile(fmt.Sprintf("%s/file_%d.txt", dir, i))
//	}
//	fmt.Println("Reading dir...")
//
//	files, _ := ioutil.ReadDir(dir)
//	fmt.Printf("Created %v files.\n", len(files))
//
//	fmt.Println("Cleaning up...")
//	os.RemoveAll(dir)
//}

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
