package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create a new map. map[<KEYTYPE>]<VALUETYPE>
	states := make(map[string]string)
	fmt.Println(states) // Empty output

	// Add items to map
	states["PA"] = "Pennsylvania"
	states["NY"] = "New York"
	states["NJ"] = "New Jersey"
	states["ME"] = "Maine"
	fmt.Println(states)

	// Accessing items in map
	pennsylvania := states["PA"]
	fmt.Println(pennsylvania)

	// Deleting items from map
	delete(states, "ME")
	fmt.Println(states)

	// Loop through map using 'for'
	for key, value := range states {
		fmt.Printf("%v: %v\n", key, value)
	}

	// List in alphabetical order - extract keys as slice then sort
	keys := make([]string, len(states))
	count := 0
	for key := range states {
		keys[count] = key
		count++
	}
	sort.Strings(keys)
	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
