package main

import (
	"github.com/JohanAanesen/CloudTech_oblig3/gofiles"
	"time"
)

func main() {

	// Get local time for now
	tempDay := time.Now().Local()

	// Format to weekday
	day := tempDay.Format("Monday")

	// Only get data if it's not the weekend (fixer.io does not update in weekends)
	if day != "Saturday" && day != "Sunday" {

		// Add currencies
		gofiles.DailyCurrencyAdder()
	}
}
