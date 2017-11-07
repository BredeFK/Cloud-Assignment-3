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
