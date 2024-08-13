package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	var a int = 1
	fmt.Println(int32(a))

	var b int32 = 2
	fmt.Println(int64(b))

	var c int64 = 3
	var d string = strconv.FormatInt(c, 10)
	fmt.Println(d)
	fmt.Println(fmt.Sprintf("%d", c))

	e, _ := strconv.Atoi(d)
	fmt.Println(e)

	f, _ := strconv.ParseInt(d, 10, 64)
	fmt.Println(f)

	var g interface{}
	g = 11

	h, ok := g.(int)
	fmt.Println(ok)
	fmt.Println(h)
	fmt.Println(g.(int))

	switch g.(type) {
	case int:
		fmt.Println(g.(int))
	case string:
		fmt.Println(g.(string))
	}

	type Student struct {
		Name string `json:"name"`
		Sex  string `json:"sex"`
		Age  int    `json:"age"`
	}

	s := Student{
		Name: "小明",
		Sex:  "男",
		Age:  11,
	}
	sByte, _ := json.Marshal(s)
	sStr := string(sByte)
	fmt.Println(sStr)

	var sP Student
	json.Unmarshal([]byte(sStr), &sP)
	fmt.Println(sP)
}
