package main

import "fmt"
import "reflect"
import "unsafe"
import "math/rand"

// the optimization of search ID

var A [1000000]int

const (
	length = 64
)

func shuffle(src []int) []int {
	dest := make([]int, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}

func minfree(a interface{}, n int) interface{} {
	// f := make([]bool, len(A))
	lista, _ := a.([]int)
	fmt.Println(reflect.TypeOf(lista))
	// for i := range lista {
	// 	if i < n {
	// 		f[i] = true
	// 	}
	// }
	// for i := 0; i < n; i++ {
	// 	if f[i] == false {
	// 		return i
	// 	}
	// }
	return lista
}

func main() {
	fmt.Println(len(A))
	fmt.Println(unsafe.Sizeof(1))
	//
	var src []int
	for i := 0; i < 10; i++ {
		src = append(src, i)
	}
	// shuffle(src)

	lista := minfree(src, 10)
	fmt.Println(reflect.TypeOf(lista))
	fmt.Println(reflect.TypeOf(src))
	// lissrc, _ := src.([]interface{})
	// fmt.Println(reflect.TypeOf())
}
