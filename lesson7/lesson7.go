package main

import "fmt"

func main() {
	// var a = [5]int{1, 2, 3, 4, 5}
	// var c = [...]string{"12", "45", "78"}
	var b = [5]int64{1: 1, 4: 4}
	b[2] = 2
	for i := 0; i < len(b); i++ {
		if i == 2 {
			fmt.Println(b[i])
		}
	}
	fmt.Println(b[2])

	for i := 0; i < len(b); i++ {
		if i == 2 {
			b[i] = 3
		}
	}
	fmt.Println(b[2])
	b[2] = 5
	fmt.Println(b[2])

	var d = [len(b) - 1]int64{}
	for i := 0; i < len(b); i++ {
		if i == 2 {
			continue
		}
		// 将元素添加到新的数组
		if i < 2 {
			d[i] = b[i]
			continue
		}
		d[i-1] = b[i-1]
	}
	fmt.Println(d[1])

	for index, value := range b {
		fmt.Printf("index-> %d, value-> %d\n", index, value)
	}

	for _, value := range b {
		fmt.Printf("value-> %d\n", value)
	}
}
