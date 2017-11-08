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
		Parameters 		map[string]string `json:"parameters"`
	} `json:"result"`
}


//Data struct
type Data struct {
	Base  string             `json:"base" bson:"base"`
	Date  string             `json:"date" bson:"date"`
	Rates map[string]float64 `json:"rates" bson:"rates"`
}