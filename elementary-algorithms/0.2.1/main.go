package main

import "fmt"
import "reflect"
import "unsafe"

// the optimization of search ID

var A [1000000]int

const (
	length = 64
)

func minfree(a interface{}) interface{} {
	return a
}

func main() {
	// fmt.Println(make([]bool, len(A)))
	fmt.Println(len(A))
	B := minfree(A)
	fmt.Println(reflect.TypeOf(B))
	fmt.Println(unsafe.Sizeof(1))
}
