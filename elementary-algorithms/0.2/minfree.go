package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

// the optimization of search ID

func shuffle(src interface{}) []int {
	sliceSrc, _ := src.([]int)
	dest := make([]int, len(sliceSrc))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(sliceSrc))
	for i, v := range perm {
		dest[v] = sliceSrc[i]
	}

	return dest
}

func minfree(a []int, n int) interface{} {
	f := make([]bool, len(a))
	for index, i := range a {
		if i < n {
			f[index] = true
		}
	}
	fmt.Println(f)
	for i := 0; i < n; i++ {
		if f[i] == false {
			return i
		}
	}
	return 0
}

func Runminfree() {
	var A []int
	fmt.Println(len(A))
	fmt.Println(unsafe.Sizeof(1))
	//
	// var src []int
	fmt.Println("shuffle----")
	// for i := 0; i < 10000; i++ {
	// 	A = append(A, i)
	// }
	// src := shuffle(A)
	A = []int{18, 4, 8, 9, 16, 1, 14, 7, 19, 3, 0, 5, 2, 11, 6}
	fmt.Println("minfree----")
	lista := minfree(A, 10)
	fmt.Println(lista)
	// fmt.Println(reflect.TypeOf(src))
	// lissrc, _ := src.([]interface{})
	// fmt.Println(reflect.TypeOf())
}
