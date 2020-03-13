package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	sha(s)

	md(s)
}

func md(s string) {
	m := md5.New()
	m.Write([]byte(s))
	ms := m.Sum(nil)
	fmt.Printf("md5: %x\n", ms)
}

func sha(s string) {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("sha1: %x\n", bs)
}
