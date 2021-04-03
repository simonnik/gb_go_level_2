package main

import (
	"fmt"
	"github.com/simonnik/go_level_1/hw2/helper"
)

func main() {
	fmt.Println("Введите сторону А:")
	a, err := helper.GetInputValue()

	if err != nil {
		return
	}
	fmt.Println("Введите сторону B:")
	b, err := helper.GetInputValue()

	if err != nil {
		return
	}
	fmt.Println("Площадь прямоугольника:", a*b)
}
