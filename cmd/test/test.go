package main

import (
	"fmt"

	"github.com/thorraythorray/go-proj/pkg/helper"
)

func main() {
	md5str, _ := helper.GetFileMD5("C:\\Users\\lei\\Downloads\\R-C.jpg")
	fmt.Println(md5str)
}
