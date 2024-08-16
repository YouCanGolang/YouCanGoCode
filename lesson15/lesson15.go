package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	for i, v := range numbers {
		fmt.Printf("index-> %d, value-> %d\n", i, v)
	}

	for i, _ := range numbers {
		fmt.Printf("index-> %d\n", i)
	}

	for _, v := range numbers {
		fmt.Printf("value-> %d\n", v)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index-> %d, value-> %d\n", i, numbers[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("第%d次循环\n", i+1)
	}

	for i := 0; i < 10; i++ {
		if i+1 == 5 {
			fmt.Printf("跳过第%d次循环\n", i+1)
			continue
		}
		fmt.Printf("第%d次循环\n", i+1)
	}

	j := 1
	for {
		fmt.Printf("死循环第%d次\n", j)
		time.Sleep(time.Duration(1) * time.Second)
		j++

		if j == 10 {
			fmt.Println("死循环退出")
			break
		}
	}

loop:
	x := 1
	for {
		fmt.Printf("死循环第%d次\n", x)
		time.Sleep(time.Duration(1) * time.Second)
		x++

		if x == 10 {
			fmt.Println("死循环从头开始执行")
			goto loop
		}
	}
}
