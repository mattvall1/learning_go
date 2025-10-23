package main

import "fmt"

func main() {
	str1 := "The quick brown fox"
	str2 := "jumped over the"
	str3 := "lazy dog"
	num1 := 51

	fmt.Println(str1, str2, str3)

	// If called, first value is length num, the second value is an error object
	strLen, err := fmt.Println("Value is: ", num1)
	// Examine error object using if stmt
	if err == nil {
		fmt.Println("Str len: ", strLen)
	}

	// Interpolation when printing
	fmt.Printf("Value of num: %v\n", num1) // `%v` = value
	fmt.Printf("Data type: %T\n", num1)    // `%T` = Type
}
