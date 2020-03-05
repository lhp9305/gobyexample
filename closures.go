package main

import "fmt"

func main() {
	netInt := intSeq()
	fmt.Println(netInt())
	fmt.Println(netInt())
	fmt.Println(netInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
