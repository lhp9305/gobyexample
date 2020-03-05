package main

import "fmt"

func main() {
	// 这里我们创建了一个字符串通道，最多允许缓存 2 个值。
	message := make(chan string, 2)
	// 由于此通道是缓冲的，因此我们可以将这些值发送到通道中
	// 而不需要相应的并发接收。
	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
