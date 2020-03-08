package main

import (
	"fmt"
	"time"
)

// 使用 go 协程和通道实现一个线程池
func main() {
	// 创建两个通道：发送任务、收集结果
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动三个worker，初始是阻塞的，因为没有传递任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 传递任务
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后，收集所有任务的返回值
	for a := 1; a <= 9; a++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
