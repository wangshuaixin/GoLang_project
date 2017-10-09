package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	content := `Invoke方法只要修改区块链的状态
就会调用 Invoke 方法。简言之，所有创建、更新和删除操作都应封装在 Invoke 方法内
因为此方法将修改区块链的状态，所以区块链 Fabric 代码会自动创建一个交易上下文，以便此方法在其中执行
对此方法的所有调用都会在区块链上记录为交易，这些交易最终被写入区块中`
	fmt.Fprintf(w, content) //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
