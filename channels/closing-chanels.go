package main

import (
	"fmt"
)

// 通道的关闭
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	// 关闭 一个通道意味着不能再向这个通道发送值了。
	//这个特性可以 用来给这个通道的接收方传达工作已经完成的信息。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		//time.Sleep(time.Second)
		jobs <- j
		fmt.Println("send job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	fmt.Println(<-done)
}
