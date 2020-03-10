package main

import "os"

//不像在有些语言中使用异常处理错误，在 Go 中则习惯通过 返回值来标示错误。
func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
