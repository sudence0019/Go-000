学习笔记

使用goroutine需要注意：

1 知道goroutine如何结束。

2 尽量让调用者去创建goroutine。

​	func **(){

​		go func(){

​		}（）

​	}

3 goroutine需要有超时检测机制，超时就关闭goroutine。

​	context

4 调用者有能力主动关闭goroutine。

​	close()

​	chan