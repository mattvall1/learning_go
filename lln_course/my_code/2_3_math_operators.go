package main

import "fmt"

// Note: No implicit conversions in Go - when using math the same type is needed
func main() {
	int1, int2, int3 := 12, 23, 34 // Can assign multiple vars in one statement
	intSum := int1 + int2 + int3
	fmt.Println("Int sum: ", intSum)

	flt1, flt2, flt3 := 12.3, 34.5, 56.7
	fltSum := flt1 + flt2 + flt3
	fmt.Println("Float sum: ", fltSum)

	// Converting data is simple. Just wrap it in the desired type
	total := flt1 + float64(int1) + flt2 + float64(int2) + flt3 + float64(int3)
	fmt.Println("Total: ", total)

	// Self-explanatory - this can be done for all operators not just multiply
	product := float64(int1) * flt3
	fmt.Println("Product: ", product)

}
