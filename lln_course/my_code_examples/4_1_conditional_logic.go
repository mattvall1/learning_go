package main

import "fmt"

func main() {
	//answer := 0
	var result string

	// Var can be set within the if/else - but is only available within that block
	if answer := 50; answer < 0 {
		result = "< 0"
	} else if answer == 0 {
		result = "= 0"
	} else {
		result = "> 0"
	}
	fmt.Println(result)
	//fmt.Println("Value of answer: ", answer) // Will only work if assigned at top

}
