package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	jsonString := `[{"name":"apple","price":2.99,"quantity":3},
				  {"name":"orange","price":1.50,"quantity":8},
				  {"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)
	fmt.Println(result)
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []CartItem {
	cartItems := make([]CartItem, 0)
	decoder := json.NewDecoder(strings.NewReader(jsonString)) // Decoder needs a reader
	_, err := decoder.Token()                                 // Guarantees delimiters are properly nested + matched
	checkError(err)

	var item CartItem
	for decoder.More() {
		err := decoder.Decode(&item)
		checkError(err)
		cartItems = append(cartItems, item)
	}

	return cartItems
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}
