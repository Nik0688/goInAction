package main

import (
	"fmt"
)

func names() (string, string) {
	return "foo", "bar"
}

func main() {
	n1, n2 := names()
	fmt.Println(n1, n2)
	fmt.Println("--------")
	n3, _ := names()
	fmt.Println(n3)
}
