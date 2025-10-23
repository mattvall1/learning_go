package main

import (
	"fmt"
	"math"
)

func main() {
	flt1, flt2, flt3 := 12.34, 34.1, 56.7
	fltSum := flt1 + flt2 + flt3
	fmt.Println("Float sum: ", fltSum) // Can have a lack of precision

	fltSum = math.Round(fltSum*100) / 100 // By using 100, some precision is lost
	fmt.Printf("Sum is now: %v\n", fltSum)

	// Pi - Geometry calculations
	fmt.Println("Value of Pi: ", math.Pi)
	radius := 15.6
	circumference := radius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference) // %.2f = two chars after decimal
}
