package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Printf("err-> %v\n", err)
		return
	}
	defer conn.Close()

	for {
		data := fmt.Sprintf("%s\n", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Printf(data)
		fmt.Fprintf(conn, data)
		time.Sleep(time.Duration(1) * time.Second)
	}
}
