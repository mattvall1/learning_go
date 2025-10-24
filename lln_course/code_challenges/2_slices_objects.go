package main

import "fmt"

func main() {
	colorNames := []string{"Red", "Green", "Blue"}
	hexValues := []int{0xFF0000, 0x00FF00, 0x0000FF}
	colors := slicesToObjects(colorNames, hexValues)
	fmt.Println(colors)
}

// slicesToObjects() returns a slice of Color objects
func slicesToObjects(colorNames []string, hexValues []int) []Color {
	// Create a slice of Color objects
	colorSlice := make([]Color, 0, len(colorNames))

	for count := range colorNames {
		colorSlice = append(colorSlice, Color{Hex: hexValues[count], Name: colorNames[count]})
		count++
	}

	return colorSlice
}

type Color struct {
	Name string
	Hex  int
}
