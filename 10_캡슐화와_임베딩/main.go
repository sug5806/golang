package main

import (
	"github.com/calender"
	"log"
)

func main() {
	date := calender.Date{}

	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(14)
	if err != nil {
		log.Fatal(err)
	}
}
