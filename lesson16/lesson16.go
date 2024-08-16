package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Printf("当前时间戳: %d\n", time.Now().Unix())
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	go printNumber()

	select {}
}

func printNumber() {
	for i := 0; i < 100; i++ {
		fmt.Printf("printNumber: %d\n", i)
		time.Sleep(time.Duration(1) * time.Second)
	}
}
