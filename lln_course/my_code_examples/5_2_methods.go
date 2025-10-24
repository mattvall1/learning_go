package main

import "fmt"

// Note: In Go, a method is a member of a type

func main() {
	cat := Cat{"Tabby", "Meow"}
	cat.Speak()
	cat.Sound = "Purr"
	cat.Speak()

	println(cat.SpeakThreeTimes())
}

// Cat Struct definition:
type Cat struct {
	Breed string
	Sound string
}

// Speak - Defining a method
func (ct Cat) Speak() {
	fmt.Printf("The %v says %v!\n", ct.Breed, ct.Sound)
}

func (ct Cat) SpeakThreeTimes() string {
	return fmt.Sprintf("%v! %v! %v!", ct.Sound, ct.Sound, ct.Sound)
}
