package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
