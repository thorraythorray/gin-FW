package main

import (
	"fmt"
)

type Person struct {
	Name    string `default:"John Doe"`
	Age     int    `default:"30"`
	Country string `default:"USA"`
}

func main() {
	p1 := Person{}
	p2 := Person{
		Name: "Alice",
	}

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}
