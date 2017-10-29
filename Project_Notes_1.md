# Project Notes
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