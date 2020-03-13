package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	err := os.Mkdir("tmp/subdir", 0755)
	check(err)

	defer os.RemoveAll("tmp/subdir")
}
