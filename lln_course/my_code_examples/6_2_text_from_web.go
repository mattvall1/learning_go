package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php" // JSON data from web

func main() {
	fmt.Println()

	// Create client object - this is an instance of a struct (params are available with this, but not needed)
	client := http.Client{}
	// Create request object
	myRequest, err := http.NewRequest("GET", url, nil)
	checkError(err)

	myRequest.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0") // This prevents 403 error

	// Using 'Do' sends the request
	myResponse, err := client.Do(myRequest)
	checkError(err)
	defer myResponse.Body.Close()

	fmt.Printf("Response type: %T\n", myResponse)

	bytes, err := io.ReadAll(myResponse.Body)
	checkError(err)
	returnedContent := string(bytes)

	fmt.Println(returnedContent)
}

// Error checking function
func checkError(err error) {
	if err != nil {
		panic(err) // `panic` stops app immedately, and displays param
	}
}
