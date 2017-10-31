package main

import "fmt"

type SortInterface interface {
	selectsort()
	bubbsort()
}
type Sorter struct {
	name string
}

func (sorter Sorter) selectsort(arry []int) {
	arraylength := len(arry)
	for i := 0; i < arraylength; i++ {
		min := i
		for j := i + 1; j < arraylength; j++ {
			if arry[j] < arry[min] {
				min = j
			}
		}
		t := arry[i]
		arry[i] = arry[min]
		arry[min] = t
	}
}

func (sorter Sorter) bubbsort(arry []int) []int {
	var temp int
	j := 1
	for j < len(arry) {
		for index := 0; index < len(arry)-j; index++ {
			if arry[index] > arry[index+1] {
				temp = arry[index]
				arry[index], arry[index+1] = arry[index+1], temp
			}
		}
		j++
	}
	return arry
}

func RunSort() {
	arry := []int{6, 2, 3, 4, 1, 5, 3, 5, 2, 4, 2, 1, 2, 3, 5, 6, 7}
	// learnsort := Sorter{name: "select sort n*n"}
	// learnsort.selectsort(arry)
	// fmt.Println(learnsort.name, arry)
	learnsort := Sorter{name: "bubble sort n*2"}
	learnsort.bubbsort(arry)
	fmt.Println(learnsort.name, arry)
}
