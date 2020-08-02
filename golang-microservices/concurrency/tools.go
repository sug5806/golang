package main

import "fmt"

func main() {
	c := make(chan string)

	fmt.Println("sending to the channel")
	c <- "hello"

	fmt.Println("receiving from the channel")
	greeting := <-c
	fmt.Println("greeting received")

	fmt.Println(greeting)

	//go helloWorld()

}

func helloWorld() {
	fmt.Println("Hello world!")
}
