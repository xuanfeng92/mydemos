package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	go func() {
		ch <- 22
		ch <- 3
		ch <- 4
		ch <- 5
		ch <- 6
		ch <- 8
		close(ch)
	}()
	time.Sleep(time.Second) //为了将ch通道填满先，计算初始的容量
	for i := 0; i < 3; i++ {
		fmt.Println("长度：", len(ch))
		fmt.Println(<-ch)
	}
	fmt.Println(<-ch)
	fmt.Println("剩余长度：", len(ch))

}
