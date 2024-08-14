package main

import "fmt"

func main() {
	printHello()

	print("我的参数内容")

	print1("1", "2")

	sumResult := sum(1, 2)
	fmt.Println(sumResult)

	fmt.Println(sum3(1))
	fmt.Println(sum3(1, 2, 3))

	fmt.Println(funcParam(sum))

	sumFunc := func(a, b int) int {
		return a + b
	}
	fmt.Println(sumFunc(1, 2))
	fmt.Println(sumFunc(3, 4))
}

func printHello() {
	fmt.Println("Hello!")
}

func print(text string) {
	fmt.Println(text)
}

func print1(text string, text2 string) {
	fmt.Println(text)
	fmt.Println(text2)
}

func sum(a, b int) int {
	return a + b
}

func sum2(a, b int) (s int) {
	s = a + b
	return s
}

func sum3(number ...int) (s int) {
	for _, n := range number {
		s += n
	}
	return
}

func funcParam(fn func(a, b int) int) int {
	a := 1
	b := 1
	return fn(a, b)
}
