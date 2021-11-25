package main

type Service struct {
	log Logger
}

type Base struct {
	Things []interface{}
}

type Person struct {
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}

type Place struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type HumanDecoder struct {
	service *Service
}

type Logger struct{}
