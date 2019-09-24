package main

import (
	"fmt"
	"time"
)

func chanelPrint(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{8, 6, 7, 5, 9, 15, 16, 0, 25}
	go chanelPrint(c)
	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}
