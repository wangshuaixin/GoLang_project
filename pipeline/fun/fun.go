package fun

import (
	"sync"
)

// fun is a pipeline conbimed with three stage

// Gen convert a list of int to a int list of channel
func Gen(nums ...int) <-chan int { //receive data from upstream
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n // send n to out
		}
		close(out)
	}()
	return out
}

// Sq receive int from another channel, and power each data, then send the data to downstream
func Sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

//Merge copy the value to downstream channel, for each channel in updtream, open a goroutines;
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// create goroutine for each channel
	// sent c to copied value untile c or done has been closed. then use wg.Done()
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// create goroutine and close it after all output has been finished.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
