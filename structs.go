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

//2d data map
type Data2d struct {
	Date string                        `json:"date" bson:"date"`
	Data map[string]map[string]float64 `json:"data" bson:"data"`
}

//Data struct
type Data struct {
	Base  string             `json:"base" bson:"base"`
	Date  string             `json:"date" bson:"date"`
	Rates map[string]float64 `json:"rates" bson:"rates"`
}
