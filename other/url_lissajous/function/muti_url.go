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
	/*
		goroutine是一种函数的并发执行方式，
		而channel是用来在goroutine之间进行参数传递。
		main函数本身也运行在一个goroutine中，
		而go function则表示创建一个新的goroutine，
		并在这个新的goroutine中执行这个函数。
	*/
	start := time.Now()
	ch := make(chan string)
	/*
		main函数中用make函数创建了一个传递string类型参数的channel，
		对每一个命令行参数，
		我们都用go这个关键字来创建一个goroutine，
		并且让函数在这个goroutine异步执行http.Get方法。
		这个程序里的io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中
	*/

	url := []string{"https://wallstreetcn.com", "https://www.apple.com", "https://www.jiqizhixin.com"}
	for _, url := range url {
		go fetch(url, ch) // start a goroutine
	}
	for range url {
		fmt.Println(<-ch) // receive from chaneel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	/*
		可以把这个Discard变量看作一个垃圾桶，可以向里面写一些不需要的数据
		因为我们需要这个方法返回的字节数，
		但是又不想要其内容。每当请求返回内容时，
		fetch函数都会往ch这个channel里写入一个字符串，
		由main函数里的第二个for循环来处理并打印channel里的这个字符串。
	*/
	defer resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("start at: %v costs: %.2fs bytes: %7d  from: %s", start, secs, nbytes, url)
}

/*
this is main
start at: 2017-09-29 22:51:54.407442 +0800 CST m=+0.001057663 costs: 0.25s bytes:   46024  from: https://www.apple.com
start at: 2017-09-29 22:51:54.407417 +0800 CST m=+0.001033030 costs: 0.25s bytes:   35072  from: https://www.jiqizhixin.com
start at: 2017-09-29 22:51:54.407435 +0800 CST m=+0.001051212 costs: 1.16s bytes:  192183  from: https://wallstreetcn.com
1.16s elapsed

可以发现在第二个网站的时候，时间比第一个还早，显然是异步
*/
