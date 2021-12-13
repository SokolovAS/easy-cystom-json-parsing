package main

type Service struct {
	log Log
}

type Base struct {
	Things []interface{} `json:"things"`
}

type Person struct {
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}

type Place struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type HumanDecode struct {
	service *Service
}

type Log struct{}
