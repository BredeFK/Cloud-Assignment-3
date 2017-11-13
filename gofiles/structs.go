package gofiles

// MongoDB struct
type MongoDB struct {
	DatabaseURL  string
	DatabaseName string
	ColCurrency  string
}

//APIPayload struct, //what we get back from Dialogflow
type APIPayload struct {
	Status struct {
		Code      int
		ErrorType string
	}
	Result struct {
		Action           *string
		ActionIncomplete bool
		Speech           string
		Parameters       map[string]interface{} `json:"parameters"`
	} `json:"result"`
}

//Data2d 2d data map
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

//Currency codes
var codes = []string{
	"AUD", "BGN", "BRL", "CAD",
	"CHF", "CNY", "CZK", "DKK",
	"EUR", "GBP", "HKD", "HRK",
	"HUF", "IDR", "ILS", "INR",
	"JPY", "KRW", "MXN", "MYR",
	"NOK", "NZD", "PHP", "PLN",
	"RON", "RUB", "SEK", "SGD",
	"THB", "TRY", "USD", "ZAR"}