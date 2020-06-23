package main

import (
	"fmt"
)

func Socialize() {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hellow!")
	fmt.Println("Nice weather, eh?")
}

func main() {
	Socialize()
}
