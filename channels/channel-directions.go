package main

import "fmt"

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// `ping` 函数定义了一个只允许发送数据的通道。尝试使用这个通
// 道来接收数据将会得到一个编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
