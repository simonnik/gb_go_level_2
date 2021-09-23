package files

import (
	"fmt"
	"os"
)

func CreateAndCloseFile(pathFile string) {
	f := createFile(pathFile)
	defer closeFile(f)
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func createFile(pathFile string) *os.File {
	f, err := os.Create(pathFile)

	if err != nil {
		panic(err)
	}

	return f
}
