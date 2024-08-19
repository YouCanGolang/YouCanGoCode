package main

import (
	"fmt"
	"strconv"
)

func main() {
	var testSlice = []int{1, 2, 3}
	fmt.Println(testSlice[4])

	a := "123t"
	num, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}
