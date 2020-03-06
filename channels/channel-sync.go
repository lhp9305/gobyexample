package main

import (
	"fmt"
	"time"
)

// 通道同步
// 我们可以使用通道来同步 Go 协程间的执行状态
func main() {
	done := make(chan bool, 1)
	go worker(done)

	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	<-done
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值通知我们已经完工了
	done <- true
}
