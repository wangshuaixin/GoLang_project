package main

import "fmt"
import "time"

func test() *int {
	x := 100
	return &x
}

func map_key() {
	m := make(map[string]string)
	m["hello"] = "echo hello"
	m["world"] = "echo world"
	m["go"] = "echo go"
	m["is"] = "echo is"
	m["cool"] = "echo cool"

	for k, v := range m {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}
}

func selecting() {

	c := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		c <- 10
	}()

	for {
		select {
		case <-c:
			println("trigger from chan C")
			return
		}
	}
}

func main() {
	println(test())

	//map
	map_key()

	//selecting channel,用途: 监听IO操作，当IO操作发生时，触发相应的动作, 与goroutine联合使用
	selecting()

	//channel
	c := make(chan int, 3)
	var send chan<- int = c
	var recv <-chan int = c
	send <- 1
	println(<-recv)
}
