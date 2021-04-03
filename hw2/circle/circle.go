package main

import (
	"fmt"
	"github.com/simonnik/go_level_1/hw2/helper"
	"math"
)

func main() {
	fmt.Println("Введите площадь круга:")
	s, err := helper.GetInputValue()

	if err != nil {
		return
	}
	fmt.Println("Длина окружности:", math.Sqrt(s*4*math.Pi))
	fmt.Println("Диаметр окружности:", 2*math.Sqrt(s/math.Pi))
}
