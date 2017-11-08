package main

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