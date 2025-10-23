package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str1, _ := reader.ReadString('\n') // Underscore ignores error
	fmt.Println(str1)

	fmt.Print("Enter number: ")
	str1, _ = reader.ReadString('\n') // Underscore ignores error. No colon as var is already declared
	floatNum1, err := strconv.ParseFloat(strings.TrimSpace(str1), 64)
	// Print error as needed
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value float: ", floatNum1)
	}

}
