package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器表示在未来某一时刻的独立事件。你告诉定时器
	// 需要等待的时间，然后它将提供一个用于通知的通道。
	// 这里的定时器将等待 2 秒。
	timer1 := time.NewTimer(time.Second * 2)
	// `<-timer1.C` 直到这个定时器的通道 `C` 明确的发送了
	// 定时器失效的值之前，将一直阻塞。
	<-timer1.C
	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()

	if stop2 := timer2.Stop(); stop2 {
		fmt.Println("timer 2 stopped")
	}

}
