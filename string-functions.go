package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	p("contains:", s.Contains("test", "es"))
	p("count:", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// 你可以在 [`strings`](http://golang.org/pkg/strings/)
	// 包文档中找到更多的函数

	// 虽然不是 `strings` 的一部分，但是仍然值得一提的是获
	// 取字符串长度和通过索引获取一个字符的机制。
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

}
