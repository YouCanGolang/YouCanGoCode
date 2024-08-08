package main

import "fmt"

func main() {
	var a []int
	a = append(a, 1)
	a = append(a, 2, 3)
	a = append(a, []int{4, 5}...)
	fmt.Println(a[0])

	var b = make([]int, len(a))
	copy(b, a)
	fmt.Println(b[0])
	fmt.Println(b[1:3])

	/*var c []string
	var d []int = make([]int, 3)
	var e []int = make([]int, 3, 4)
	var f = []int{1, 2, 3}*/
}
