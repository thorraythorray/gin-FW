package main

import "fmt"

func main() {
	var i interface{} // 定义一个空接口类型

	i = "Hello, Go!" // 将一个字符串赋值给接口

	// 断言 i 是否为字符串类型
	if s, ok := i.(string); ok {
		fmt.Println("i 是一个字符串：", s)
	} else {
		fmt.Println("i 不是一个字符串")
	}

	// 断言 i 是否为整数类型
	if _, ok := i.(int); ok {
		fmt.Println("i 是一个整数")
	} else {
		fmt.Println("i 不是一个整数")
	}
	type FormHandler interface {
		Validate() error
	}
	var f FormHandler
	fmt.Printf("%T", f)
}
