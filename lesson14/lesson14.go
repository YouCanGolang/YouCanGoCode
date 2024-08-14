package main

import "fmt"

func main() {
	a := 1
	b := 2

	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a < b || a == b")
	}

	if a > b {
		fmt.Println("a > b")
	} else if a == b {
		fmt.Println("a == b")
	} else {
		fmt.Println("a < b")
	}

	switch a {
	case 1:
		fmt.Println("a == 1")
		fallthrough
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	default:
		fmt.Println("a value null")
	}
}
