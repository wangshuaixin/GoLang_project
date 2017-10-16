# GoLang_project

This repo is used for collecting some practice and fancy trick for Golang.

---
用来记载Golang练习，



---

### List

+ sorter 快速和冒泡排序/rapid sequencing and Bubble Sort
+ crypto 简单加密算法Golang/simple cryptography in GO
+ go-python go语言和python相互调用/ contact between Golang and Python
+ spider 爬虫/ a web crawler in Golang
+ web 运用http协议搭建web / build simple website based on golang net package
+ blockchain 从零开始搭建blockchain / build a blockchain from scratch

#### Dash means extend feature based on initial program/ ‘-’ 意味着其他的简单扩展



---


### LOG
+  The article [How to write Go code](https://golang.org/doc/code.html) is essential to novice.
+  run Go code:
    For `build` or `install` You need to have your files in package directory. and `build <directory name>`
    
    For `go run`, you need to supply all files as argument:
    `go run *.go`

0. 注意 install之后是放在`$HOME/go/bin/`的位置
1. 大写小写字母开头的类型，变量和函数等表示可见性。
2. 数组切片 {slice, ...type} 作为syntactic sugar
3. 先写变量名，再写类型名，[see reason for](#Fancy Golang)
4. package 用来组织，只有一个main
5. 内置关键字为小写
6. 具有闭包和[匿名函数](http://books.studygolang.com/gopl-zh/ch5/ch5-06.html)，虽然其他语言也有
7. defer 栈 [see this blog](https://chinazt.cc/2017/06/30/deferde-shi-yong-gui-ze/)
8. :=
9. 值传递和指针传递
10. 获取变量类型两种方法，Sprintf("%T",a),reflect.Typeof()
11. interface{} 对标object
12. 基于指针对象的方法要写在函数名前[see](http://books.studygolang.com/gopl-zh/ch6/ch6-02.html)
13. 写python或者其他静态类型语言就是总是觉得编译不好，在同一个package里面相互调用的时候一般要在一起编译 go build 1.go 2.go. 然后./main 分级的话则可以相互调用？ 
14. 注意包的位置是代表整个dir
15. 关于规范化 https://github.com/golang/go/wiki/CodeReviewComments 
16. 每个可见方法都要有注释
17. VScode 的golint总是不太灵光，自动导入import包还是要少用
18. 翻转slice [...]inst{0,1,2,3,4,5}
19. cap 取到满足条件的2,4,8,16等值
20. 判断slice里面元素对比，只有bytes.Equal 有单独的方法，其他没有
21. slice元素是间接引用，slice的值在不同时刻可能包含不同的元素，因为底层数组会被修改。
22. 看 一个map！
```
userAttrList := map[string]map[string]map[int]bool{
    "first": {
        "second": {
            3: false,
        },
    },
}
```
23. 单元测试感觉很不错





## [Fancy Golang](#FancyGolang)
1. https://blog.golang.org/gos-declaration-syntax
