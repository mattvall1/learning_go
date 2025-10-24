package main

import (
	"fmt"
	"strconv"
)

func main() {
	value1 := "10"
	value2 := "5.5"
	operation := "+"
	result := calculate(value1, value2, operation)
	fmt.Println(result)
}

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	var result float64
	flt1 := convertInputToValue(input1)
	flt2 := convertInputToValue(input2)
	switch {
	case operation == "+":
		result = addValues(flt1, flt2)
	case operation == "-":
		result = subtractValues(flt1, flt2)
	case operation == "*":
		result = multiplyValues(flt1, flt2)
	case operation == "/":
		result = divideValues(flt1, flt2)
	default:
		result = 0
	}

	return result
}

func convertInputToValue(input string) float64 {
	float, err := strconv.ParseFloat(input, 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number. ", input)
		panic(message)
	}
	return float
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
