package main

import (
	"fmt"
	"time"
)

func main() {
	weekday := time.Now().Weekday()
	fmt.Printf("Today is: %v\n", weekday)

	dayNum := int(weekday)
	fmt.Printf("Day as number is: %v\n", dayNum)

	var result string
	//dayNum = 7 // Test
	switch dayNum { // switch dayNum = 7; dayNum {
	case 1:
		result = "It's Monday."
	case 2:
		result = "It's Tuesday."
	case 3:
		result = "It's Wednesday."
	case 4:
		result = "It's Thursday."
	case 5:
		result = "It's Friday."
	default:
		result = "It's the Weekend."
	}
	fmt.Println(result)

	// No need for break statement. 'fallthrough' = after any code, the next case will still execute
	aNumber := 6
	switch {
	case aNumber > 0:
		result = "More than 0"
		fallthrough
	case aNumber == 0:
		result = "Equals 0"
		fallthrough
	default:
		result = "Less than 0"
	}
	fmt.Println(result)
}
