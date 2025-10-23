package main

import (
	"fmt"
	"time"
)

func main() {
	tme := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at: %s\n", tme)

	now := time.Now()
	fmt.Printf("Time now: %s\n", now)
	fmt.Printf("This object type is: %T\n", now)

	fmt.Printf(now.Format(time.ANSIC) + "\n") // Format a time object

	// Add one day
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf(tomorrow.Format(time.ANSIC) + "\n") // Format a time object

	// Set own specific format
	format := "Mon 2006-02-01" // Details on this date: https://stackoverflow.com/questions/42217308
	fmt.Printf(tomorrow.Format(format) + "\n")
}
