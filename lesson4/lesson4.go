package main

import "fmt"

func main() {
	var A int = 100
	var B int = 10

	fmt.Printf("A + B = %d\n", A+B)
	fmt.Printf("A - B = %d\n", A-B)
	fmt.Printf("A * B = %d\n", A*B)
	fmt.Printf("A / B = %d\n", A/B)
	fmt.Printf("A 与 B 求余 = %d\n", A%B)

	A--
	fmt.Printf("A++ 结果为 %d\n", A)

	B--
	fmt.Printf("B-- 结果为 %d\n", B)

	fmt.Printf("A == B 结果为 %v\n", A == B)
	fmt.Printf("A != B 结果为 %v\n", A == B)
	fmt.Printf("A > B 结果为 %v\n", A == B)
	fmt.Printf("A < B 结果为 %v\n", A == B)
	fmt.Printf("A >= B 结果为 %v\n", A == B)
	fmt.Printf("A <= B 结果为 %v\n", A == B)

	var C bool = true
	var D bool = false

	fmt.Printf("C && D 结果为 %v\n", C && D)
	fmt.Printf("C || D 结果为 %v\n", C || D)
	fmt.Printf("!C 结果为 %v\n", !C)
}
