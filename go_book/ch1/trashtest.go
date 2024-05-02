package main

import (
	"fmt"
	u "github.com/bcicen/go-units"
)

func main() {
	// convert a simple float from celsius to fahrenheit
	fmt.Println(u.MustConvertFloat(25.5, u.Celsius, u.Fahrenheit)) // outputs "77.9 fahrenheit"

	// convert a units.Value instance
	val := u.NewValue(25.5, u.Celsius)

	fmt.Println(val)                           // "25.5 celsius"
	fmt.Println(val.MustConvert(u.Fahrenheit)) // "77.9 fahrenheit"
	fmt.Println(val.MustConvert(u.Kelvin))     // "298.65 kelvin"
}
