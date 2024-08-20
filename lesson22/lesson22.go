package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	deferPanic()

	go func() {
		for {
			fmt.Println("Running...")
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	select {}
}

func deferPanic() {
	fmt.Printf("函数进入...\n")

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover err-> %v\n", err)
		}
	}()

	fmt.Printf("函数执行...\n")

	panic(errors.New("执行出错"))

	fmt.Printf("函数return...\n")

	return
}
