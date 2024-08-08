package main

import "fmt"

func main() {
	var a map[string]int
	a = make(map[string]int)
	a["A"] = 1
	fmt.Println(a["A"])
	delete(a, "A")
	fmt.Println(a["A"])

	ok, value := a["A"]
	fmt.Println(ok)
	fmt.Println(value)

	var b = make(map[string]int)
	b["A"] = 1

	var c = map[string]int{
		"A": 1,
		"B": 2,
	}
	for key, value := range c {
		fmt.Printf("key-> %s, value-> %d\n", key, value)
	}
}
