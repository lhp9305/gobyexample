package main

import (
	"fmt"
	"time"
)

// Go 中可以通过 Go 协程、通道和打点器来支持了速率限制。
func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	// 这个 `limiter` 通道将每 200ms 接收一个值。这个是
	// 速率限制任务中的管理器。
	limiter := time.Tick(time.Millisecond * 200)

	// 通过在每次请求前阻塞 `limiter` 通道的一个接收，我们限制
	// 自己每 200ms 执行一次请求。
	for req := range requests {
		<-limiter
		fmt.Println("requests", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 有时候我们想临时进行速率限制，并且不影响整体的速率控制，
	// 我们可以通过[通道缓冲](channel-buffering.html)来实现。
	// 这个 `burstyLimiter` 通道用来进行 3 次临时的脉冲型速率限制。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
