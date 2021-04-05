package main

import (
	"fmt"
	"github.com/simonnik/go_level_1/hw2/helper"
)

func main() {
	fmt.Println("Введите число a:")
	a, err := helper.GetInputValue()

	if err != nil {
		return
	}

	fmt.Println("Введите число b:")
	b, err := helper.GetInputValue()

	if err != nil {
		return
	}

	fmt.Println("Выберите операцию: +, -, *, /")

	var operation string

	_, err = fmt.Scanln(&operation)

	if err != nil {
		return
	}

	switch operation {
	case "+":
		fmt.Println("Результат:", a+b)
	case "-":
		fmt.Println("Результат:", a-b)
	case "*":
		fmt.Println("Результат:", a*b)
	case "/":
		if b != 0 {
			fmt.Println("Результат:", a/b)
		} else {
			fmt.Println("Деление на 0")
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}
