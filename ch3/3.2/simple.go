package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside goroutine")
	go func() {
		fmt.Println("inside gouroutine")
	}()
	fmt.Println("outside again")
	runtime.Gosched()
}
