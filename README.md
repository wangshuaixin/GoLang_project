# GoLang_project

This repo is used for collecting some practice and fancy trick for Golang.

---
用来记载GOlang练习和一些语法，


### 语法
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
13. 写python或者其他静态类型语言就是总是觉得编译不好，在同一个package里面相互调用的时候一般要在一起编译 go build . 然后./main 分级的话则可以相互调用？ 



---
#### sorter 是一个快排和冒泡的区别算法。／sorter is a procedure-oriented way for sorting algorithm

#### GO_graphic is a graphic algorithm for Go

---


## [Fancy Golang](#FancyGolang)
1. https://blog.golang.org/gos-declaration-syntax
