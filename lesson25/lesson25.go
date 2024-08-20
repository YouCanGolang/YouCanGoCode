package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入时间戳 (输入 exit 退出)")

	for {
		fmt.Printf("时间戳: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入错误: ", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("退出程序")
			break
		}

		ts, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("输入的时间戳格式有误")
			continue
		}

		date := time.Unix(ts, 64).Format("2006-01-02 15:04:05")
		fmt.Println("转换结果: ", date)
	}
}
