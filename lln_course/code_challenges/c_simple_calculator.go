package main

import (
	"fmt"
	"strconv"
)

func main() {
	value1 := "10"
	value2 := "5.5"
	result := calculate(value1, value2)
	fmt.Println(result)
}

// calculate() returns the the result of adding 2 numbers after parsing them from strings
func calculate(value1 string, value2 string) float64 {
	// Convert strings to float64
	val1f, err1 := strconv.ParseFloat(value1, 64)
	if err1 != nil {
		fmt.Println(err1)
	}
	val2f, err2 := strconv.ParseFloat(value2, 64)
	if err2 != nil {
		fmt.Println(err2)
	}

	// Calculate and return
	return val1f + val2f
}
