// Package main

/*
课后练习 2.1
------------

将练习 1.2 中的生产者消费者模型修改成为多个生产者和多个消费者模式
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	id       int // 任务序号
	producer int // 是哪个生产者发的消息
	delay    int // 消费者处理这个消息需要多少秒
}

func main() {
	// 创建一个 channel 用于收发消息
	ch := make(chan Task, 10)
	// 计数器，100 个任务，处理完就结束
	// 读写 remainTasks 变量的时候，要用互斥锁
	taskId := 0
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	// 10 个 Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(producer int) {
			fmt.Printf("Producer #%d start.\n", producer)
			for {
				sProduct := rand.Intn(5) + 1 // 生产消息需要的秒数
				sConsume := rand.Intn(5) + 1 // 消费消息需要的秒数
				time.Sleep(time.Duration(sProduct) * time.Second)
				if taskId >= 100 {
					break
				}
				lock.Lock()
				taskId += 1
				task := Task{id: taskId, producer: producer, delay: sConsume}
				lock.Unlock()
				fmt.Printf("Produce Task(id=%d, producer=%d)\n", task.id, producer)
				ch <- task
			}
			fmt.Printf("Producer #%d exit.\n", producer)
			wg.Done()
		}(i)
	}
	// 10 个 Consumer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(consumer int) {
			fmt.Printf("Consumer #%d start.\n", consumer)
			for task := range ch {
				fmt.Printf("Start Task(id=%d, producer=%d) by consumer#%d\n",
					task.id, task.producer, consumer)
				if task.id == 100 {
					fmt.Println("All Tasks is received, close channel...")
					close(ch)
				}
				time.Sleep(time.Duration(task.delay) * time.Second)
				fmt.Printf("Finish Task(id=%d, producer=%d) by consumer#%d\n",
					task.id, task.producer, consumer)
			}
			fmt.Printf("Consumer #%d exit.\n", consumer)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
