### Project Notes
由于之间的readme太长，所以用一个markdown来记录新的笔记。

+ 关于类型转换，和接口查询。
    interface{}是一个通用类型，可以储存任意类型的值。声明了一个名为general的interface{}值，然后依次存储一个浮点数和一个整数，通过switch general.(type)判断general存储的值的类型，然后按照类型使用类似general.(int)的方法对general进行类型转换（实际上是接口查询或者接口转换，作用是判断general是否实现了int的方法；之所以叫类型转换，是因为interface{}里面没有定义任何方法，而且它的作用也的确跟类型转换一样）。在这里用数字做例子，是因为数字简单，容易理解。我们知道，数字可以进行四则运算，在这里，我对general转换得到的数字进行加法运算，加2,然后打印加法的结果。从运行结果来看，经过转换后的数字可以进行正确的数字运算。
    
    var general interface{}
    general = 6.6
    reflect.TypeOf(general)
    general = 2
    reflect.TypeOf(general)


+ make 虽然用得多与定义变量，但是不能在函数外定义
+ copy函数不能将slice变array， 如果要实现这种操作则：
    var src []int 
	copy(A[:], src) //A is [10000]int, is a array. So that use [:] to convert to slice 
+ 万能类型(interface{})很神奇，就像 C 里面的 void*，但是C本身是一门不安全的语言，可以直接操纵原始的二进制位，所以 void* 是有必要的，但是这个东西对于强类型的Go是非常有害的和不安全的，它让你失去了静态强类型所带来的好处，很多本该在编译期就检查出来的类型问题，变成了运行时错误，我用过一些数据库方面的第三方库，就因为其内部使用了大量 interface{} 导致程序跑起来时不时有运行时类型错误非常头疼，另外大量使用 type assertion 代码是非常丑陋的。所以一定要尽量避免使用这个类型，这也是我泡 #go-nuts 频道时里面的大神给的建议

+ 关于类型转换的一些要素
    一般的强制类型转换 T(x) ， 对于非常量的x来说，需要满足下面的：
    x is assignable to T. 
    x's type and T have identical underlying types.
    x's type and T are unnamed pointer types and their pointer base types have identical underlying types.
    x's type and T are both integer or floating point types.
    x's type and T are both complex types.
    x is an integer or a slice of bytes or runes and T is a string type.
    x is a string and T is a slice of bytes or runes.

+ 静态类型是在声明的时候赋予的.interface变量区别于动态类型（一种在运行的时候多种类型的聚合体）。比如（The dynamic type may vary during execution but values stored in interface variables are always assignable to the static type of the variable.）

+ type assertion: `variable.(Type)`

+ 如果存在类似于main函数调用父目录的接口，则可以使用导入 `improt ".." `这样的方法
+ 通道(channel)和映射(map)与切片(slice)一样，也是引用类型


#### section2
---
1. 非常推荐使用 WaitGroup 来 跟踪 goroutine 的工作是否完成。WaitGroup 是一个计数信号量，我们可以利用它来统计所有的 goroutine 是不是都完成了工作
    + 每个 goroutine 完成其工作后，就会递减 WaitGroup 变量的 计数值，当这个值递减到 0 时，我们就知道所有的工作都做完了

2. 有了闭包，函数可以直接访问到那些没有作为参数传入的变量。匿名函数并没有拿到这些变量的副本，而是直接访问外层函数作用域中声明的这些变量本身。
    + 对于常规的函数，随着外部变量的循环或者改变，除非使用函数参数的形式传值给函数，否则大部分的goroutine最终会使用同一个值，很可能就是循环的slice最后一个值。

3. 方法声明中的值和值的指针的区别：
```
    // 方法声明为使用 defaultMatcher 类型的值作为接收者
    func (m defaultMatcher) Search(feed *Feed, searchTerm string)
    // 声明一个指向 defaultMatcher 类型值的指针 dm := new(defaultMatch)
    // 编译器会解开 dm 指针的引用，使用对应的值调用方法 dm.Search(feed, "test")
    // 方法声明为使用指向 defaultMatcher 类型值的指针作为接收者
    func (m *defaultMatcher) Search(feed *Feed, searchTerm string)
    // 声明一个 defaultMatcher 类型的值 var dm defaultMatch
    // 编译器会自动生成指针引用 dm 值，使用指针调用方法 dm.Search(feed, "test")
```

+ 因为大部分方法在被调用后都需要维护接收者的值的状态，所以，一个最佳实践是，将方法 的接收者声明为指针。对于 defaultMatcher 类型来说，使用值作为接收者是因为创建一个 defaultMatcher 类型的值不需要分配内存。由于 defaultMatcher 不需要维护状态，所以 不需要指针形式的接收者。
+ 与直接通过值或者指针调用方法不同，如果通过接口类型的值调用方法，规则有很大不同；使用指针作为接收者声明的方法，只能在接口类型的值是一个指针的时 候被调用。使用值作为接收者声明的方法，在接口类型的值为值或者指针时，都可以被调用。

4. 接口方法调用所受限制的例子：

```
    // 方法声明为使用指向 defaultMatcher 类型值的指针作为接收者
    func (m *defaultMatcher) Search(feed *Feed, searchTerm string)
    // 通过 interface 类型的值来调用方法
    var dm defaultMatcher
    var matcher Matcher = dm // 将值赋值给接口类型 matcher.Search(feed, "test") // 使用值来调用接口方法
    > go build
    cannot use dm (type defaultMatcher) as type Matcher in assignment
    // 方法声明为使用 defaultMatcher 类型的值作为接收者
    func (m defaultMatcher) Search(feed *Feed, searchTerm string)
    // 通过 interface 类型的值来调用方法
    var dm defaultMatcher
    var matcher Matcher = &dm // 将指针赋值给接口类型 matcher.Search(feed, "test") // 使用指针来调用接口方法
        > go build
        Build Successful
```

+ 除了 Search 方法，defaultMatcher 类型不需要为实现接口做更多的事情了。从这段代 码之后，不论是 defaultMatcher 类型的值还是指针，都满足 Matcher 接口，都可以作为 Matcher 类型的值使用。这是代码可以工作的关键。defaultMatcher 类型的值和指针现在还 可以作为 Matcher 的值，赋值或者传递给接受 Matcher 类型值的函数。

+ 这里提到了使用指针等等的问题：一个函数何时该用指针类型做receiver对初学者而言一直是个头疼的问题。如果不知道该如何取舍，选择指针类型的receiver。但有些时候value receiver更加合适，比如对象是一些轻量级的不变的structs，使用value receiver会更加高效。下面是列举了一些常用的判断指导。

        1、如果receiver是map、func或者chan，不要使用指针
        2、如果receiver是slice并且该函数并不会修改此slice，不要使用指针
        3、如果该函数会修改receiver，此时一定要用指针
        4、如果receiver是struct并且包含互斥类型sync.Mutex，或者是类似的同步变量，receiver必须是指针，这样可以避免对象拷贝
        5、如果receiver是较大的struct或者array，使用指针则更加高效。多大才算大？假设struct内所有成员都要作为函数变量传进去，如果觉得这时数据太多，就是struct太大
        6、如果receiver是struct，array或者slice，并且其中某个element指向了某个可变量，则这个时候receiver选指针会使代码的意图更加明显
        7、如果receiver使较小的struct或者array，并且其变量都是些不变量、常量，例如time.Time，value receiver更加适合，因为value receiver可以减少需要回收的垃圾量。
        8、最后，如果不确定用哪个，使用指针类的receiver

+ 在golang中没有明确的shuffle方法，如果要生成随机数组的话，需要使用rand.Perm 如果要每次都生成随机，则使用rand.NewSource()
+ python 的sorted是快排？golang里面的？
+ golang里面的切片类似与数组，但是数组知道确定的值的含量。指针数组
+ 数组是构造切片和映射的基石。
+ Go 语言里切片经常用来处理数据的集合，映射用来处理具有键值对结构的数据。
+ 内置函数 make 可以创建切片和映射，并指定原始的长度和容量。也可以直接使用切片和映射字面量，或者使用字面量作为变量的初始值。
+ 切片有容量限制，不过可以使用内置的 append 函数扩展容量。
+ 映射的增长没有容量或者任何限制。
+ 内置函数 len 可以用来获取切片或者映射的长度。
+ 内置函数 cap 只能用于切片。
+ 通过组合，可以创建多维数组和多维切片。也可以使用切片或者其他映射作为映射的值。
但是切片不能用作映射的键。
+ 将切片或者映射传递给函数成本很小，并且不会复制底层的数据结构。