package helper

import (
	"fmt"
	"strconv"
)

func GetInputValue() (float64, error) {
	var s string
	_, errScan := fmt.Scan(&s)

	if errScan != nil {
		fmt.Println("Введите число")
		return 0, errScan
	}

	n, err := strconv.ParseFloat(s, 8)

	if err != nil {
		fmt.Println("Введите число")
	}

	return n, err
}
