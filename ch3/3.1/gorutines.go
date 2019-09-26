package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout)
	time.Sleep(time.Second * 30)
	fmt.Println("timed out")
	os.Exit(0)

}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}
