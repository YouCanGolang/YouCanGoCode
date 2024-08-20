package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("handleConnection, remote-> %s\n", conn.RemoteAddr().String())
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("断开连接\n")
				break
			}
			fmt.Printf("err-> %v", err)
		}
		fmt.Printf("接收数据: %s", data)
	}
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Printf("err-> %v\n", err)
		return
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("err-> %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
