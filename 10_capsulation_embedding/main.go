package main

import (
	"fmt"
	"golang/10_capsulation_embedding/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())

}
