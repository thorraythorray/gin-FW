package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Address struct {
	Location string
}

type User struct {
	Name    string
	Age     int
	Address `mapstructure:",squash"`
}

func Test(options ...User) {
	for _, v := range options {
		fmt.Println(v)
	}
}

func main() {
	data := map[string]interface{}{
		"Name":     "John",
		"Age":      30,
		"Location": "New York",
	}

	var user User

	// 使用 mapstructure.Decode() 函数将数据映射到结构体
	err := mapstructure.Decode(data, &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Age: %d\n", user.Age)
	fmt.Printf("Location: %s\n", user.Location)
}
