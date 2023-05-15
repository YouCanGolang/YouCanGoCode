package main

import "fmt"

func main() {
	book := "You Can Golang！"
	fmt.Printf("%v\n", book)
	fmt.Printf("%s\n", book)

	var a = 10.2
	var b = 10.2345
	var c = 10
	fmt.Printf("a: %f\n", a)
	fmt.Printf("b: %0.2f\n", b)
	fmt.Printf("c: %d\n", c)

	fmt.Printf("You\n Can\n GoLang\n")
}
