package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("Foo", "1")
	fmt.Println("Foo:", os.Getenv("Foo"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], "=", pair[1])
	}
}
