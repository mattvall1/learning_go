package main

import "fmt"

// Note: Structs = similar to objects in other languages

func main() {
	britishShortHair := Cat{"English Short Hair", 4.5}
	fmt.Printf("%+v\n", britishShortHair) // `+v` = names + values of fields
	fmt.Printf("Breed: %v, Weight (KG): %v", britishShortHair.Breed, britishShortHair.WeightKG)

	// Change value within struct
	britishShortHair.WeightKG = 4.8
	fmt.Printf("Breed: %v, Weight (KG): %v", britishShortHair.Breed, britishShortHair.WeightKG)
}

// Cat Struct definition:
type Cat struct {
	Breed    string
	WeightKG float64
}
