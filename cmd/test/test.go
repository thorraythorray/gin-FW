package main

import "fmt"

// 定义一个接口
type Shape interface {
	Area() float64
}

// 定义一个结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// 结构体Rectangle实现接口Shape的Area方法
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// 创建一个结构体指针
	rectanglePtr := Rectangle{
		Width:  10,
		Height: 5,
	}

	// 将结构体指针赋值给空接口
	var s Shape = &rectanglePtr

	// 调用接口方法
	fmt.Println("Area:", s.Area()) // 输出：Area: 50
}
