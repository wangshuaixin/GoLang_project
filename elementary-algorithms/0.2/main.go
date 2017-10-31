package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

type Key struct {
	xs int
}

func partition(key *Key) int {
	return key.xs + 1
}

func main() {
	k := Key()
	result := partition()
	fmt.Println(result)
}
