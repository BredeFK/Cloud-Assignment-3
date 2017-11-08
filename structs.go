package main

// MongoDB struct
type MongoDB struct {
	DatabaseURL  string
	DatabaseName string
	ColCurrency  string
}

//ApiPayload struct, //what we get back from dialogflow
type ApiPayload struct {
	Status struct {
		Code      int
		ErrorType string
	}
	Result struct {
		Action           *string
		ActionIncomplete bool
		Speech           string
		Parameters       map[string]string `json:"parameters"`
	} `json:"result"`
}

//Data2d struct 2d array
type Data2d struct {
	Date string                        `json:"date"`
	Data map[string]map[string]float64 `json:"data"`
}

// DataDate struct struct for date
type DataDate struct {
	Date string
	Map  Data2d
}

//Data struct
type Data struct {
	Base  string             `json:"base" bson:"base"`
	Date  string             `json:"date" bson:"date"`
	Rates map[string]float64 `json:"rates" bson:"rates"`
}
