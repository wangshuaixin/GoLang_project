package main

import (
	"fmt"
	"strings"
)

func getName(params ...interface{}) {
	var paramSlice []string
	for _, param := range params {
		paramSlice = append(paramSlice, param.(string))
	}
	aa := strings.Join(paramSlice, "_") // Join 方法第2个参数是 string 而不是 rune
	fmt.Println(aa)
}

func main() {
	getName("redis", "100", "master")
}
