package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "tim", age: 18})

	fmt.Println(person{name: "tim"})

	fmt.Println(&person{name: "tom"})

	s := person{name: "sean", age: 10}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp)
	fmt.Println(sp.name)

	sp.age = 11
	fmt.Println("sp:", sp, &sp)
	fmt.Println("s:", s, &s)
}
