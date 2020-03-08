package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// 原子计数器
func main() {
	var ops uint64 = 0

	// 启动 50 个协程对计数器进行累加
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 使用 `AddUint64` 来让计数器自动增加，使用
				// `&` 语法来给出 `ops` 的内存地址。
				atomic.AddUint64(&ops, 1)
				// 允许其它 Go 协程的执行
				runtime.Gosched()
			}
		}()
	}
	// 等待一秒，让 ops 的自加操作执行一会
	time.Sleep(time.Second)
	// 为了在计数器还在被其它 Go 协程更新时，安全的使用它，
	// 我们通过 `LoadUint64` 将当前值的拷贝提取到 `opsFinal`
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
