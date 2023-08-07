package main

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}
