package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("unix", "../socketfile")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(conn, "Hello")
}
