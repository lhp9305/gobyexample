package main

import "fmt"

func main() {
	// 通道 (Channels) 是连接多个 Go 协程的管道。
	// 你可以从一个 Go 协程 将值发送到通道，然后在别的 Go 协程中接收。

	message := make(chan string)

	go func() {
		message <- "ping"
	}()

	msg := <-message
	fmt.Println(msg)
}
