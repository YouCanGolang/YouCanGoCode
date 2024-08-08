package main

import "fmt"

func main() {
	type xueKe struct {
		ShuXue int
		YuWen  int
	}

	type student struct {
		Name  string
		Sex   string
		Age   int
		XueKe xueKe
		xueKe
	}

	var st student
	st.Name = "小明"
	st.Sex = "男"
	st.Age = 11
	st.XueKe.ShuXue = 91
	st.XueKe.YuWen = 90
	st.ShuXue = 91
	st.YuWen = 90
	fmt.Printf("%#v\n", st)

	st1 := student{
		Name: "小红",
		Sex:  "女",
		Age:  1,
		XueKe: xueKe{
			ShuXue: 80,
			YuWen:  81,
		},
		xueKe: xueKe{
			ShuXue: 90,
		},
	}
	fmt.Printf("%#v\n", st1)

	/*st1 := new(student)
	var st2 *student
	st3 := &student{}
	var stSlice []student*/
}
