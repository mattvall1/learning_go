package main

import "fmt"

// Note: Array in Go is an ordered collection of values or references
//       It also has a fixed number of items (size is declared, then unchangeable after)
//		 All items in array must be same type

func main() {
	var colours [3]string // Declaring `colours` as a collection of 3 string values
	colours[0] = "Pink"
	colours[1] = "Purple"
	colours[2] = "Red"
	fmt.Println(colours)    // Outputs '[Pink Purple Red]'
	fmt.Println(colours[1]) // Outputs '[Pink Purple Red] \n Purple'

	var nums = [5]int{5, 9, 7, 1, 3} // Declaring + assigning values in-line
	fmt.Println(nums)
	fmt.Println("Num of colours:", len(colours))
	fmt.Println("Num of nums:", len(nums))
}
