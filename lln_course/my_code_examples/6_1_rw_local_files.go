package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./from_string.txt"
	file, err := os.Create(fileName)
	// 'defer' waits until everthing in function has executed, then executes this
	defer file.Close()
	checkError(err)
	fileLength, err := io.WriteString(file, "Writing to the file from Go")
	checkError(err)
	fmt.Printf("Wrote file with %v chars\n", fileLength)

	readFile(fileName)
}

// Error checking functions are good practice
func checkError(err error) {
	if err != nil {
		panic(err) // `panic` stops app immedately, and displays param
	}
}

func readFile(fileName string) {
	// Data is returned as slice of bytes
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:\n", string(data)) // string() displays data as string not bytes
}
