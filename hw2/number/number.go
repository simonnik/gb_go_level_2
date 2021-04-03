package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Введите число 0-999:")
	n, err := getInputValue()

	if err != nil {
		return
	}

	hundreds := n / 100
	tens := (n / 10) % 10
	units := n % 10
	fmt.Println("Длина окружности:", hundreds, tens, units)
}

func getInputValue() (int64, error) {
	var s string
	_, errScan := fmt.Scan(&s)

	if errScan != nil {
		fmt.Println("Введите число")
		return 0, errScan
	}

	n, err := strconv.ParseInt(s, 10, 64)

	if err != nil || (n < 0 || n > 999) {
		fmt.Println("Введите число 0-999")
		err = errors.New("wrong number")
	}

	return n, err
}
