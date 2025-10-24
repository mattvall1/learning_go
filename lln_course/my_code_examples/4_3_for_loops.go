package main

import "fmt"

// Note: While loops do not exist in Go, they can be replaced with this.

func main() {
	// Create colours map as in example 3_3
	var colours = []string{"Pink", "Purple", "Red", "Blue", "Green"}

	// Conventional iterative loop that watches a var
	for i := 0; i < len(colours); i++ {
		println(colours[i])
	}

	// Range loop (assigning index `i`) - same output as above
	for i := range colours {
		println(colours[i])
	}

	// _ means ignore this variable - same output as above
	for _, colour := range colours {
		println(colour)
	}

	// Create states map as in example 3_4
	states := make(map[string]string)
	states["PA"] = "Pennsylvania"
	states["NY"] = "New York"
	states["NJ"] = "New Jersey"
	states["ME"] = "Maine"

	// Ignore value in this case (the _ is redundant - this is an example)
	for state, _ := range states {
		println(states[state])
	}

	// Equivalent of while loop - Loop using boolean expression
	value := 0
	sum := 0
	for value < 5 {
		sum += value
		fmt.Printf("Value: %v\n", value)
		fmt.Printf("Sum: %v\n", sum)
		value++
	}

	// Loop using goto
	sum = 1
	for sum < 1000 {
		sum += sum
		if sum > 200 {
			goto theEnd
		}
	}
	// theEnd lable - by calling `goto`, this code is run
theEnd:
	println("End of program")
	fmt.Printf("Sum: %v\n", sum)
}
