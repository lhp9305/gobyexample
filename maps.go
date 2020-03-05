package main

import "fmt"

func main() {
	// 要创建一个空 map，需要使用内建的 `make`:
	// `make(map[key-type]val-type)`.
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("del:", m)

	// 当从一个 map 中取值时，可选的第二返回值指示这个键
	// 是否在这个 map 中。这可以用来消除键不存在和键有零值，
	// 像 `0` 或者 `""` 而产生的歧义。这里我们不需要值，所以
	// 用_空白标识符(blank identifier)_忽略。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
