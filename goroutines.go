package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 同步调用
	f("direct")

	// 使用 go 关键字启动协程去执行函数调用
	go f("goroutine")

	// 也可以使用匿名函数启动一个 go 协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
