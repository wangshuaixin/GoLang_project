#### tips

多个函数可以从相同的channel读取直到channel被关闭，这被叫做fan-out，提供一种方法在一组worker中平衡地分配CPU和IO的利用。

若一个函数可以从多个输入读取数据然后处理直到所有的多个输入的channel关闭，这个被叫做fan-in。

我们通过运行从相同的输入channel中读取数据的两个实例sq来改变我们的管道，使用merge

refer: http://www.opscoder.info/golang_pipeline.html