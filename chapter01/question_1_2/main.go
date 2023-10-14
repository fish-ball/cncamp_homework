// Package main

/*
课后练习 1.2
------------

基于 Channel 编写一个简单的单线程生产者消费者模型：

队列：队列长度 10，队列元素类型为 int
生产者：每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	ticker := time.NewTicker(time.Second)
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ticker.C:
				fmt.Println("in:", i)
				ch <- i
			}
		}
		close(ch)
	}()
	for x := range ch {
		time.Sleep(time.Second)
		fmt.Println("out:", x)
	}
}
