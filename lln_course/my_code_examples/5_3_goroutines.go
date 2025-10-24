package main

import (
	"fmt"
	"time"
)

func main() {
	// Using 'go' turns a function into a goroutine - This runs in background while main() continues
	go say("Hello from Goroutine")
	fmt.Println("Hello from main()")

	// Can create functions with no name (anonymous), immediately following 'go' keyword
	go func(message string) {
		fmt.Println(message)
	}("Hello from anonymous function.")

	time.Sleep(5 * time.Second)
	fmt.Println("End")
}

func say(message string) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}
