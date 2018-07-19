package main

/*多路复合计算*/
import (
	"fmt"
	"time"
)

func do_stuff(x int) int {
	time.Sleep(time.Second * 1)
	return 100 - x
}

func branch(x int) chan int {
	ch := make(chan int)
	go func() {
		ch <- do_stuff(x)
	}()
	return ch
}

func fanIn(chs ...chan int) chan int {
	ch := make(chan int)
	for _, c := range chs {
		go func(c chan int) {
			ch <- <-c
		}(c)
	}
	return ch
}

func main() {
	result := fanIn(branch(1), branch(2), branch(3))
	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
