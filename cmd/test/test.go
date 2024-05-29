package main

import (
	"fmt"

	"github.com/thorraythorray/go-Jarvis/utils"
)

func main() {
	a, _ := utils.GetHostname()
	fmt.Println(a)
}
