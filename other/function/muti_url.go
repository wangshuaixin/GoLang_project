package function

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// MutiURL : use goroutine and channel to get plantiful url
func MutiURL() {
	start := time.Now()
	ch := make(chan string)

	url := "https://www.baidu.com"
	go fetch(url, ch) // start a goroutine
	fmt.Println(<-ch)
	fmt.Printf("%.2fs elapsed\n,", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
