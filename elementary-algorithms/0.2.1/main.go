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

func shuffle(src interface{}) interface{} {
	// listsrc, _ := src.([]int)
	src, _ = src.([]int)
	// fmt.Println(reflect.TypeOf(listsrc))
	fmt.Println(reflect.TypeOf(src))
	dest := make([]int, len(src))

	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}

func minfree(a interface{}, n int) interface{} {
	f := make([]bool, len(A))
	lista, _ := a.([]int)
	fmt.Println(reflect.TypeOf(lista))
	for i := range lista {
		if i < n {
			f[i] = true
		}
	}
	for i := 0; i < n; i++ {
		if f[i] == false {
			return i
		}
	}
	return lista
}

func main() {
	fmt.Println(len(A))
	fmt.Println(unsafe.Sizeof(1))
	//
	// var src []int
	fmt.Println("dash----")
	for i := 0; i < 1000000; i++ {
		A[i] = i
	}
	src := shuffle(A)
	fmt.Println(reflect.TypeOf(src))

	lista := minfree(A, 10)
	fmt.Println(reflect.TypeOf(lista))
	fmt.Println(lista)
	// fmt.Println(reflect.TypeOf(src))
	// lissrc, _ := src.([]interface{})
	// fmt.Println(reflect.TypeOf())
}
