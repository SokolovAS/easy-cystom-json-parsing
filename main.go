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

type HumanDecoder struct {
	service Service
}

func NewHumanDecoder(service Service) *HumanDecoder {
	return &HumanDecoder{service: service}
}

func NewService(log Logger) *Service {
	return &Service{log: log}
}

type ByAge []Person
type ByCity []Place

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a ByCity) Len() int           { return len(a) }
func (a ByCity) Less(i, j int) bool { return len(a[i].City) < len(a[j].City) }
func (a ByCity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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

type Logging interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Logger struct{}

func (s *Logger) Println(v ...interface{}) {
	fmt.Println(v)
}

func (s *Logger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v)
}

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

var jsonStr = []byte(`
{
    "things": [
        {
            "name": "Alice",
            "age": 37
        },
        {
            "city": "Ipoh",
            "country": "Malaysia"
        },
        {
            "name": "Bob",
            "age": 36
        },
        {
            "city": "Northampton",
            "country": "England"
        },
 		{
            "name": "Albert",
            "age": 3
        },
		{
            "city": "Dnipro",
            "country": "Ukraine"
        },
		{
            "name": "Roman",
            "age": 32
        },
		{
            "city": "New York City",
            "country": "US"
        }
    ]
}`)

func showResults(data []byte) {
	logger := Logger{}
	s := Service{logger}
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
