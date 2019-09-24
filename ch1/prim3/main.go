package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	con, _ := net.Dial("tcp", "golang.org:80")
	fmt.Fprintf(con, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(con).ReadString('\n')
	fmt.Println(status)
}
