// Package temperature implements logic to convert the temperature from different metric systems.
package temperature

import "fmt"

// The CelsiusToFahrenheit function accepts temperature value in Celsius and
// returns an appropriate value in Fahrenheit.
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	tC := 9.0

	fmt.Println("Current temperature in Fahrenheit", CelsiusToFahrenheit(tC))
}
