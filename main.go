package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

type HumanDecoding interface {
	Decode(data []byte) ([]Person, []Place)
}

type Logging interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

func NewHumanDecoder(service *Service) *HumanDecoder {
	return &HumanDecoder{service: service}
}

func NewService(log Logger) *Service {
	return &Service{log: log}
}

func (h *HumanDecoder) Decode(data []byte) ([]Person, []Place) {
	var b Base
	var persons []Person
	var places []Place

	err := json.Unmarshal(data, &b)
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range b.Things {
		data := val.(map[string]interface{})
		if data["name"] != nil {
			person := Person{data["name"].(string), data["age"].(float64)}
			persons = append(persons, person)
		} else {
			place := Place{data["city"].(string), data["country"].(string)}
			places = append(places, place)
		}
	}
	return persons, places
}

func (s *Logger) Println(v ...interface{}) {
	fmt.Println(v)
}

func (s *Logger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v)
}

func showResults(data []byte) {
	logger := Logger{}
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
