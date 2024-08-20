package main

import (
	"errors"
	"fmt"
)

func main() {
	deferPrint()

	deferPanic()
}

func deferPrint() {
	fmt.Printf("函数进入...\n")

	defer fmt.Println("函数结束，defer逻辑调用...")

	fmt.Printf("函数执行...\n")

	fmt.Printf("函数return...\n")

	return
}

func deferPanic() {
	fmt.Printf("函数进入...\n")

	defer fmt.Println("函数结束，defer逻辑调用...")

	fmt.Printf("函数执行...\n")

	panic(errors.New("执行出错"))

	fmt.Printf("函数return...\n")

	return
}
