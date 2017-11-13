### Project notes 4

1. 
一种变量的命名形式

	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.goinggo.net/feeds/posts/default?alt=rss",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}

2. 常用的BYTE 数据类型以无差别字节流的形式存储任何种类的二进制数据。二进制数据通常由数字化的信息（如，电子表格、程序装入模块和数字化声音模式等等）组成。

BYTE 数据类型不具有最大大小。BYTE 列具有 2 个字节的理论限制和磁盘容量确定的实际限制。
可以存储、检索、更新或删除 BYTE 列的内容。但是，不能在算术或字符串运算中使用 BYTE 操作数，也不能使用 UPDATE 语句的 SET 子句将文字指定给 BYTE 列。

主要是不需要数据的转换

