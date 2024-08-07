package main

import "fmt"

func main() {
	var a int = 1
	var b *int
	var c *int
	b = &a
	c = &a

	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(*c)

	a = 2
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(*c)
}
