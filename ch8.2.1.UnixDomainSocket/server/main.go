package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("unix", "../socketfile")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		// エラー処理
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Accept:%v\n", conn.LocalAddr())
	buffer := make([]byte, 100)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received: %s\n", string(buffer[:length]))

}
