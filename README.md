# GoLang_project

This repo is used for collecting some practice and fancy trick for Golang.

---
用来记载Golang练习，



---

### List
+ [project notes](https://github.com/dyllanwli/GoLang_project/blob/master/Project_Notes_1.md)用来记录golang的feature
+ [elementary-algorithms](https://github.com/dyllanwli/GoLang_project/tree/master/elementary-algorithms) is a golang implementation algorithm
+ [pipeline in golang](https://github.com/dyllanwli/GoLang_project/tree/master/pipeline)
+ [graphic](https://github.com/dyllanwli/GoLang_project/tree/master/graphic)
+ [other](https://github.com/dyllanwli/GoLang_project/tree/master/other) 快速和冒泡排序 CHANNEL goroutine defer之类的用法/rapid sequencing and Bubble Sort and channel goroutin defer
+ [crypto](https://github.com/dyllanwli/GoLang_project/tree/master/crypto) 简单加密算法Golang/simple cryptography in GO
+ [go-python](https://github.com/dyllanwli/GoLang_project/tree/master/go-python) go语言和python相互调用/ contact between Golang and Python
+ [spider](https://github.com/dyllanwli/GoLang_project/tree/master/spider) 爬虫/ a web crawler in Golang
+ [web](https://github.com/dyllanwli/GoLang_project/tree/master/web) 运用http协议搭建web / build simple website based on golang net package
+ [blockchain](https://github.com/dyllanwli/GoLang_project/tree/master/blockchain) 从零开始搭建blockchain / build a blockchain from scratch

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

24. printf
```
package main
import "fmt"
import "os"
type point struct {
    x, y int
}
func main() {
//Go 为常规 Go 值的格式化设计提供了多种打印方式。例如，这里打印了 point 结构体的一个实例。
    p := point{1, 2}
    fmt.Printf("%v\n", p) // {1 2}
//如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
    fmt.Printf("%+v\n", p) // {x:1 y:2}
//%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
    fmt.Printf("%#v\n", p) // main.point{x:1, y:2}
//需要打印值的类型，使用 %T。
    fmt.Printf("%T\n", p) // main.point
//格式化布尔值是简单的。
    fmt.Printf("%t\n", true)
//格式化整形数有多种方式，使用 %d进行标准的十进制格式化。
    fmt.Printf("%d\n", 123)
//这个输出二进制表示形式。
    fmt.Printf("%b\n", 14)
这个输出给定整数的对应字符。
    fmt.Printf("%c\n", 33)
%x 提供十六进制编码。
    fmt.Printf("%x\n", 456)
//对于浮点型同样有很多的格式化选项。使用 %f 进行最基本的十进制格式化。
    fmt.Printf("%f\n", 78.9)
//%e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式。
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)
//使用 %s 进行基本的字符串输出。
    fmt.Printf("%s\n", "\"string\"")
//像 Go 源代码中那样带有双引号的输出，使用 %q。
    fmt.Printf("%q\n", "\"string\"")
//和上面的整形数一样，%x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示。
    fmt.Printf("%x\n", "hex this")
//要输出一个指针的值，使用 %p。
    fmt.Printf("%p\n", &p)
//当输出数字的时候，你将经常想要控制输出结果的宽度和精度，可以使用在 % 后面使用数字来控制输出宽度。默认结果使用右对齐并且通过空格来填充空白部分。
    fmt.Printf("|%6d|%6d|\n", 12, 345)
//你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度。
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
//要最对齐，使用 - 标志。
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
//你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。这是基本的右对齐宽度表示。
    fmt.Printf("|%6s|%6s|\n", "foo", "b")
//要左对齐，和数字一样，使用 - 标志。
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
//到目前为止，我们已经看过 Printf了，它通过 os.Stdout输出格式化的字符串。Sprintf 则格式化并返回一个字符串而不带任何输出。
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)
//你可以使用 Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout。
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
```

25. golang 里面的时间 1e+09 = time.Second 

    //比如
    time.Sleep(1e+09) = time.Sleep(time.Second)

26. underline before import a package, like `_ "github.com/..."` 

the short answer is that it is for importing a package solely for its side -effects.

27. 在Go中对于管道没有正式的定义，它只是很多并发程序中的一种。一般说来，一个管道就是一系列通过channel连接的stages，每个stage都是运行着相同函数的goroutines的集合，在每个stage中，goroutine干着下面三件事儿：

    通过channel从上游接收数据；
    对这些数据执行一些函数，通常再产生新的数据；
    再把数据发送到下游的channel中。
    每个stage都有一些上下游channel，除了最开始和最后的stage，最开始的stage通常叫做source或是producer，最后的stage叫做sink或是consumer。

28. can not range interface 的解决：https://stackoverflow.com/questions/42054248/cannot-range-over-list-type-interface-in-function-using-go

29. 感觉需要单独简历一个markdown来记录这些tricky了。新建一个吧。


## [Fancy Golang](#FancyGolang)
1. https://blog.golang.org/gos-declaration-syntax
