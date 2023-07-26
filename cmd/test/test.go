package main

import (
	"fmt"
	"math"
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

// func main() {
// 	data := map[string]interface{}{
// 		"Name":     "John",
// 		"Age":      30,
// 		"Location": "New York",
// 	}

// 	var user User

// 	// 使用 mapstructure.Decode() 函数将数据映射到结构体
// 	err := mapstructure.Decode(data, &user)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

//		fmt.Printf("Name: %s\n", user.Name)
//		fmt.Printf("Age: %d\n", user.Age)
//		fmt.Printf("Location: %s\n", user.Location)
//	}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

// func main() {
// 	rectangle := Rectangle{Width: 5, Height: 3}
// 	circle := Circle{Radius: 2}

// 	PrintArea(rectangle) // 输出：Area: 15.00
// 	PrintArea(circle)    // 输出：Area: 12.57
// }

// 定义一个名为 "Person" 的结构体
type Person struct {
	Name string
	Age  int
}

// 函数接收一个名为 "p"，类型为 "Person" 的结构体作为参数
func PrintPerson(p Person) {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

// func main() {
// 	// 创建一个 Person 结构体变量并赋值
// 	person := Person{
// 		Name: "Alice",
// 		Age:  30,
// 	}

// 	// 调用函数并将结构体变量作为参数传递
// 	PrintPerson(person)
// }

type InnerStruct struct {
	Field int
}

func (is InnerStruct) Method() {
	fmt.Println("InnerStruct Method called with Field:", is.Field)
}

type OuterStruct struct {
	InnerStruct
	OuterField string
}

func (os OuterStruct) Method() {
	fmt.Println("OuterStruct Method called with OuterField:", os.OuterField)
}

type Person1 struct {
	Name string
	Age  int
}

type Person2 struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	// 定义两个结构体变量
	person1 := Person1{Name: "Alice", Age: 30}
	person2 := Person2{}

	// 快速赋值：将 person1 的字段值复制给 person2
	person2 = person1

	// 输出两个结构体的字段值
	fmt.Println("Person 1:", person1) // 输出: Person 1: {Alice 30}
	fmt.Println("Person 2:", person2) // 输出: Person 2: {Alice 30}
}
