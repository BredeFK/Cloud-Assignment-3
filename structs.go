package main

// MongoDB struct
type MongoDB struct {
	DatabaseURL  string
	DatabaseName string
	ColCurrency  string
}

// Currency struct
type Currency struct {
	Base  map[string]float64  `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}