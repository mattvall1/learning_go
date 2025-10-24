package main

import "fmt"

func main() {
	doSomething()
}

// No in/out function
func doSomething() {
	fmt.Println("Do something")

	val1 := 5
	val2 := 6
	fmt.Printf("The sum of %v and %v is %v\n", val1, val2, addValues(val1, val2))

	val3 := 9
	fmt.Printf("The sum of the values is %v\n", addAllValues(val1, val2, val3))

	val4 := 10
	val5 := 70
	sum, count, avg := addAllValuesMultiReturn(val1, val2, val3, val4, val5)
	fmt.Printf("The sum of %v values is %v which averages to be %v\n", count, sum, avg)
}

// Go is strictly typed, but you only need to declare type once if they're all the same.
func addValues(val1, val2 int) int {
	return val1 + val2
}

// Declaring like this means the function receives a slice of ints
func addAllValues(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

// In Go, multiple values can be returned from one function
func addAllValuesMultiReturn(values ...int) (int, int, float64) {
	sum := 0
	for _, val := range values {
		sum += val
	}
	count := len(values)
	avg := float64(sum) / float64(count)
	return sum, count, avg
}
