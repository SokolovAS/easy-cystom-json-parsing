package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

type HumanDecoder interface {
	Decode(data []byte) ([]Person, []Place)
}

type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

func NewHumanDecoder(service *Service) *HumanDecode {
	return &HumanDecode{service: service}
}

func NewService(log Log) *Service {
	return &Service{log: log}
}

func (h *HumanDecode) Decode(data []byte) ([]Person, []Place) {
	var (
		b       Base
		persons []Person
		places  []Place
	)

	if err := json.Unmarshal(data, &b); err != nil {
		log.Fatal(err)
	}

	for _, val := range b.Things {
		data := val.(map[string]interface{})
		if data["name"] != nil {
			name, ok := data["name"].(string)
			checkError("Error type assertion 'name'", ok)
			age, ok := data["age"].(float64)
			checkError("Error type assertion 'age'", ok)
			person := Person{name, age}
			persons = append(persons, person)
			continue
		}
		city, ok := data["city"].(string)
		checkError("Error type assertion 'city'", ok)
		country, ok := data["country"].(string)
		checkError("Error type assertion 'country'", ok)
		place := Place{city, country}
		places = append(places, place)

	}
	return persons, places
}

func checkError(s string, ok bool) {
	if !ok {
		log.Fatal("Error type assertion 'country'")
	}
}

func (s *Log) Println(v ...interface{}) {
	fmt.Println(v)
}

func (s *Log) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v)
}

func showResults(data []byte) {
	logger := Log{}
	s := NewService(logger)
	h := NewHumanDecoder(s)
	persons, places := h.Decode(data)

	sort.Sort(ByAge(persons))
	sort.Sort(ByCity(places))

	h.service.log.Println(persons)
	h.service.log.Println(places)
}

func main() {
	showResults(jsonStr)
}
