package main

import (
	"fmt"
	"sort"
)

// Note: In general - its recommended to use slices instead of arrays to manage ordered data

func main() {
	// No number in len = slice (not a memory-efficient way of creating slices)
	var colours = []string{"Pink", "Purple", "Red", "Blue", "Green"}
	fmt.Println(colours)

	// Memory efficient slice | 0 = num initial items, 5 = capacity
	var colors = make([]string, 0, 5) // `make()` - declares item and allocates memory
	colors = append(colors, "Pink", "Purple", "Red", "Blue", "Green")
	fmt.Println(colors)
	// When more items are appended - slice is automatically recreated in background
	colors = append(colors, "Magenta", "Orange")
	fmt.Println(colors)

	// Removing items from slice using custom function
	colors = remove(colors, 3) // Remove 3rd item in slice
	fmt.Println(colors)

	// Sorting a slice (alphabetical sort ascending)
	sort.Strings(colors)
	fmt.Println(colors)
}

// Removing items from slice is more complicated, hence new function `remove()`
func remove(slice []string, index int) []string { // Return type before curly bracket
	// Starting with index position + 1, return everything else
	return append(slice[:index], slice[index+1:]...)
}
