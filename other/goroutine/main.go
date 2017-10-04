// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"time"
// )

// var counter int = 0

// // 通过共享内存的方式来进行goroutine, 开发效率较为低
// func Count(lock *sync.Mutex) {
// 	lock.Lock()
// 	counter++ //将内存锁住
// 	fmt.Println(counter, time.Now())
// 	lock.Unlock()
// }

// func main() {
// 	lock := &sync.Mutex{}
// 	for i := 0; i < 10; i++ {
// 		go Count(lock)
// 	}

// 	for {
// 		lock.Lock()
// 		c := counter
// 		lock.Unlock()
// 		runtime.Gosched()

// 		if c >= 10 {
// 			break
// 		}
// 	}
// }

package main

import (
	"fmt"
)

func Count(ch chan int) {
	fmt.Println("Counting", ch)
	ch <- 1
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}

}
