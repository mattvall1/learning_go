package main

import "fmt"

func main() {
	int1 := 25
	// var point *int // Declaring var with no value. * = deference operator - points to a var of type int
	var point *int = &int1 // Points to memory address of int1

	if point == nil {
		fmt.Println("point is nil")
	} else {
		fmt.Println("Value of p:", *point)
	}

	// Same can be achieved with other data types
	float1 := 25.15
	point1 := &float1 // Infer memory address with &
	fmt.Println("Value 1: ", *point1)

	// Modifying pointer thru pointer
	*point1 = *point1 / 13
	fmt.Println("Pointer 1: ", *point1)
	// When pointers are changed that are pointing to var, the var is changing
	fmt.Println("Original value 1: ", float1)

}
