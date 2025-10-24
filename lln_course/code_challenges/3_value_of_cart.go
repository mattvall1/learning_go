package main

import (
	"fmt"
)

func main() {
	var cart []CartItem
	var apples = CartItem{"apple", 2.99, 3}
	var oranges = CartItem{"orange", 1.50, 8}
	var bananas = CartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)

	fmt.Println(result)
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []CartItem) float64 {
	var result float64

	for _, cartItem := range cart {
		result += cartItem.Price * float64(cartItem.Quantity)
	}
	return result
	//return math.Round(result*100) / 100 // Two trailing 0s = round to 2dp
}

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}
