package main

import (
	"fmt"

	"./fun"
)

func main() {
	//construct the pipeline
	c := fun.Gen(2, 3)
	out := fun.Sq(c)

	// purchase output
	fmt.Println(<-out)
	fmt.Println(<-out)

	for n := range fun.Sq(fun.Gen(2, 3)) {
		fmt.Println(n)
	}
	fmt.Println("-----")
	//use merge
	in := fun.Gen(2, 3)

	// 从两个goroutine中分配sq的工作
	c1 := fun.Sq(in)
	c2 := fun.Sq(in)

	// 从c1 c2消费merge之后的output
	for n := range fun.Merge(c1, c2) {
		fmt.Println(n) // 4 然后 9, 或是 9 然后 4
	}
}
