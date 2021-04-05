package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Введите число N:")
	n, err := getInputValue()

	if err != nil {
		return
	}

	simples := getSimpleNumbers(n)
	for _, p := range simples {
		fmt.Println(p)
	}
}

func getSimpleNumbers(N int) (primes []int) {
	b := make([]bool, N)
	primes = append(primes, 0)
	primes = append(primes, 1)

	for i := 2; i < N; i++ {
		if b[i] {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}

	return
}

func getInputValue() (int, error) {
	var s string
	_, errScan := fmt.Scan(&s)

	if errScan != nil {
		fmt.Println("Введите число")
		return 0, errScan
	}

	n, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		err = errors.New("wrong number")
	}

	return int(n), err
}
