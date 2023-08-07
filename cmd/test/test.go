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

func main() {
	// 定义一个存放 Person 结构体的 slice
	persons := []Person{
		{
			Name: "Alice",
			Age:  30,
			Address: Address{
				City:    "New York",
				Country: "USA",
			},
		},
		{
			Name: "Bob",
			Age:  25,
			Address: Address{
				City:    "London",
				Country: "UK",
			},
		},
	}

	// 打印 slice 中的每个 Person 结构体

}
