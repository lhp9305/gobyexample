package main

import "fmt"

// 个人理解：
// * 用来声明指针 (*int ) 或者 用来取指针指向的值
// & 取地址符

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)

	// 打印指针地址
	fmt.Println("pointer:", &i)
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}
