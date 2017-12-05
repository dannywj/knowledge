package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		/*
			select 默认是阻塞的，只有当监听的 channel 中有发送或接收可以进行时才会运行，当多
			个 channel 都准备好的时候，select 是随机的选择一个执行的。
			在select 里面还有 default 语法，select 其实就是类似 switch 的功能，default 就是当监听的channel 都没有准备好的时候，默认执行的（select 不再阻塞等待 channel）。

			如下示例是监听2个chan，哪个chan有值进来就执行哪个。在main中quit是循环结束后才进来的，也就是做为结束标志来输出
		*/
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

/*
1
1
2
3
5
8
13
21
34
55
quit
*/
