package main

import (
	"errors"
	"fmt"
	"strconv"
)

type CheckErr struct {
	Code int
	Msg  string
}

func main() {
	ok, err := checkNum(9)
	if err != nil {
		fmt.Printf("检查错误信息: %v\n", err)
	}
	fmt.Printf("检查结果: %v\n", ok)

	ok1, err1 := checkNum(17)
	if err1 != nil {
		fmt.Printf("检查错误信息1: %v\n", err1)
	}
	fmt.Printf("检查结果1: %v\n", ok1)

	if _, err2 := checkNum(18); err != nil {
		fmt.Printf("检查错误信息2: %v\n", err2)
	}

	if _, checkErr := checkNum1(9); checkErr != nil {
		fmt.Printf("检查错误信息3: errCode-> %d, errMesg-> %s\n", checkErr.Code, checkErr.Msg)
	}

	a := "1234t"
	num, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		fmt.Printf("转换出错: err-> %v, str-> %s\n", err, err.Error())
		return
	}
	fmt.Printf("转换成功, num-> %d\n", num)
}

func checkNum(num int) (bool, error) {
	if num < 10 {
		return false, fmt.Errorf("数字太小, 输入数字为: %d", num)
	}

	if num > 15 {
		return false, errors.New(fmt.Sprintf("数字太大, 输入数字为: %d", num))
	}
	return true, nil
}

func checkNum1(num int) (bool, *CheckErr) {
	if num < 10 {
		return false, &CheckErr{
			Code: 1,
			Msg:  fmt.Sprintf("数字太小了, 输入的数字是: %d", num),
		}
	}
	return true, nil
}
