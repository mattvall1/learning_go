package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php" // JSON data from web

func main() {
	content := readContent()
	tours := toursFromJSON(content)
	for _, tour := range tours {
		price, _ := strconv.ParseFloat(tour.Price, 16)
		fmt.Printf("%v: $%v\n", tour.Name, price)
	}
}

// This is a near-copy of 6_2 except the data is returned
func readContent() string {
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

	bytes, err := io.ReadAll(myResponse.Body)
	checkError(err)
	returnedContent := string(bytes)

	return returnedContent
}

func toursFromJSON(content string) []Tour {
	tours := make([]Tour, 0)
	decoder := json.NewDecoder(strings.NewReader(content)) // Decoder needs a reader
	_, err := decoder.Token()                              // Guarantees delimiters are properly nested + matched
	checkError(err)

	var tour Tour
	for decoder.More() { // 'More' looks for next item in JSON content, if found it returns true and vice versa
		err := decoder.Decode(&tour) // Pass memory address of Tour object. This reads, parses and adds data to matching fields in 'Tour' struct
		checkError(err)
		tours = append(tours, tour)
	}
	return tours
}

// Error checking function
func checkError(err error) {
	if err != nil {
		panic(err) // `panic` stops app immedately, and displays param
	}
}

// Tour struct
type Tour struct {
	Name, Price string // Data can be different types, but in this example the price is wrapped in quotes
}
