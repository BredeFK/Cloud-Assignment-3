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

		// Get currencies from fixer.io
		data2d := gofiles.GetCurrency()

		// Set up db
		db := gofiles.SetupDB()

		// Add currencies to mongodb
		db.DailyCurrencyAdder(data2d)
	}
}
