### note2

+ *内置类型*是由语言提供的一组类型。比如数值类型、字符串类型 和布尔类型。
+ Go 语言里的引用类型有如下几个:切片、映射、通道、接口和函数类型
+ 当声明上述类型的变量时，创建的变量被称作标头(header)值。从技术细节上说，字符串也是一种引用类型。 每个引用类型创建的标头值是包含一个指向底层数据结构的指针。每个引用类型还包含一组独特的字段，用于管理底层数据结构。因为标头值是为复制而设计的，所以永远不需要共享一个引用类型的值。标头值里包含一个指针，因此通过复制来传递一个引用类型的值的副本，本质上就是在共享底层数据结构。
+ *结构类型*可以用来描述一组数据值，这组值的本质即可以是原始的，也可以是非原始的。如 果决定在某些东西需要删除或者添加某个结构类型的值时该结构类型的值不应该被更改，那么需 要遵守之前提到的内置类型和引用类型的规范。
+ 多态是指代码可以根据类型的具体实现采取不同行为的能力。如果一个类型实现了某个接口，所有使用这个接口的地方，都可以支持这种类型的值。


### 接口实现

接口是用来定义行为的类型。这些被定义的行为不由接口直接实现，而是通过方法由用户定义的类型实现。如果用户定义的类型实现了某个接口类型声明的一组方法，那么这个用户定义的类型的值就可以赋给这个接口类型的值。这个赋值会把用户定义的类型的值存入接口类型的值。

所以对于接口值的调用会执行接口值里面的存储的用户定义的类型的值对应的方法。任何用户定义的类型都可以实现任何接口，所以对接口值方法的调用自然就是一种多态。在这个关系里面，用户定义的类型通常叫做*实体类型*，原因是如果离开内部存储的用户定义的类型的值的实现，接口值没有具体的行为。

### 方法集

方法集是定义接口的接受规则。

方法集定义了一组关联到给定类型的值或者指针的方法。定义方法时使用的接收者的类型决定了这个方法是关联到值，还是关联到指针，还是两个都关联。

- - - 
|Values| Methods Receivers |
|-|-|
|T |(t T)|
|*T |(t T) and (t *T)|

T类型的值只包含值接受者声明的方法。而指向T类似的指针的方法集既包含值接受者声明的方法，也包含指针接受者声明的方法。从值的角度看这些规则，会显得复杂很多。

下面从接收者上看

|Methods Receivers |Values |
|-|-|
|(t T) |T and *T |
|(t *T) |*T|

这个规则说，如果使用指针接收者来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口。如果使用值 接收者来实现一个接口，那么那个类型的值和指针都能够实现对应的接口。

则有时候需要传入地址（&），有时候不需要


+ 使用关键字 struct 或者通过指定已经存在的类型，可以声明用户定义的类型。
+ 方法提供了一种给用户定义的类型增加行为的方式。
+ 设计类型时需要确认类型的本质是原始的，还是非原始的。
+ 接口是声明了一组行为并支持多态的类型。
+ 嵌入类型提供了扩展类型的能力，而无需使用继承。
+ 标识符要么是从包里公开的，要么是在包里未公开的。


### 并发和并行

并发(concurrency)不是并行(parallelism)。并行是让不同的代码片段同时在不同的物理处 理器上执行。并行的关键是同时做很多事情，而并发是指同时管理很多事情，这些事情可能只做了一半就被暂停去做别的事情了。在很多情况下，并发的效果比并行好，因为操作系统和硬件的总资源一般很少，但能支持系统同时做很多事情。这种“使用较少的资源做更多的事情”的哲学，也是指导 Go 语言设计的哲学。

WaitGroup 是一个计数信号量，可以用来记录并维护运行的 goroutine。如果 WaitGroup 的值大于 0，Wait 方法就会阻塞。将这个 WaitGroup 的值设置为 2，表示有两个正在运行的 goroutine。为了减小 WaitGroup 的值并最终释放 main 函数，使用 defer 声明在函数退出时 调用 Done 方法。

#### basis 58 算法
Bitcoin uses the Base58 algorithm to convert public keys into human readable format. The algorithm is very similar to famous Base64, but it uses shorter alphabet: some letters were removed from the alphabet to avoid some attacks that use letters similarity. Thus, there are no these symbols: 0 (zero), O (capital o), I (capital i), l (lowercase L), because they look similar. Also, there are no + and / symbols.

### 竞争机制

如果两个或者多个 goroutine 在没有互相同步的情况下，访问某个共享的资源，并试图同时读和写这个资源，就处于相互竞争的状态，这种情况被称作竞争状态(race candition)。竞争状态 的存在是让并发程序变得复杂的地方，十分容易引起潜在问题。对一个共享资源的读和写操作必 须是原子化的，换句话说，同一时刻只能有一个 goroutine 对共享资源进行读和写操作。

一种修正代码、消除竞争状态的办法是，使用 Go 语言提供的锁机制，来锁住共享资源，从而保证 goroutine 的同步状态。


### 锁住共享资源

Go 语言提供了传统的同步 goroutine 的机制，就是对共享资源加锁。如果需要顺序访问一个 整型变量或者一段代码，atomic 和 sync 包里的函数提供了很好的解决方案。下面我们了解一 下 atomic 包里的几个函数以及 sync 包里的 mutex 类型。

1. 下面是具有atomic进行强制添加：

另外两个有用的原子函数是 LoadInt64 和 StoreInt64

    // incCounter increments the package level counter variable.
    func incCounter(id int) {
        // Schedule the call to Done to tell main we are done.
        defer wg.Done()

        for count := 0; count < 2; count++ {
            // Safely Add One To Counter.
            atomic.AddInt64(&counter, 1)

            // Yield the thread and be placed back in queue.
            runtime.Gosched()
        }
    }

如果没有保护则：

    func incCounter(id int) {
        // Schedule the call to Done to tell main we are done.
        defer wg.Done()

        for count := 0; count < 2; count++ {
            // Capture the value of Counter.
            value := counter
            // Yield the thread and be placed back in queue.
            runtime.Gosched()
            // Increment our local value of Counter.
            value++
            // Store the value back into Counter.
            counter = value
        }
    }

这样的。


2. 还可以通过互斥锁来进行共享资源访问。

同步访问共享资源的方式是使用互斥锁(mutex)。互斥锁这个名字来自互斥(mutual exclusion)的概念。互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以 执行这个临界区代码。

	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Only allow one goroutine through this
		// critical section at a time.
		mutex.Lock()
		// Capture the value of counter.
		value := counter

		// Yield the thread and be placed back in queue.
		runtime.Gosched()

		// Increment our local value of counter.
		value++

		// Store the value back into counter.
		counter = value

		mutex.Unlock()
		// Release the lock and allow any
		// waiting goroutine through.
	}

### channel 
原子函数和互斥锁都能工作，但是依靠它们都不会让编写并发程序变得更简单，更不容易出 错，或者更有趣。在 Go 语言里，你不仅可以使用原子函数和互斥锁来保证对共享资源的安全访 问以及消除竞争状态，还可以使用通道，通过发送和接收需要共享的资源，在 goroutine 之间做同步。

当一个资源需要在 goroutine 之间共享时，通道在 goroutine 之间架起了一个管道，并提供了 确保同步交换数据的机制。声明通道时，需要指定将要被共享的数据的类型。可以通过通道共享 内置类型、命名类型、结构类型和引用类型的值或者指针。

+ 在只有一个CPU工作的时候，永远是最后一个先执行。剩下的按顺序执行