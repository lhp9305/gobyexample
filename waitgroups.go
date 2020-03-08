package main

import (
	"fmt"
	"sync"
	"time"
)

// 协程同步
func main() {
	// 这个 WaitGroup 被用于等待该函数开启的所有协程。
	var wg sync.WaitGroup

	// 开启几个协程，并为其递增 WaitGroup 的计数器。
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	// 阻塞，直到 WaitGroup 计数器恢复为 0，即所有协程的工作都已经完成。
	wg.Wait()
}

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)

	wg.Done()
}
