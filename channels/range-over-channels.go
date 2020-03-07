package main

import "fmt"

// 通道遍历
// 我们也可以使用 range 来遍历从通道中取得的值。
func main() {
	queue := make(chan int, 2)

	go func() {
		for i := 0; i < 9; i++ {
			queue <- i
		}
		close(queue)
	}()

	for elem := range queue {
		fmt.Println(elem)
	}
}
