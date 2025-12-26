package main

import (
	"fmt"
	"strings"
)

func main() {
	// 定义 string 字符串

	var str string = "你好golang"
	s := "你好 go"
	fmt.Println(str, s)

	// 转义字符
	// \n 换行符

	s1 := "hello\nworld"
	fmt.Println(s1)

	// 输出 \\ 反斜杠
	s2 := "~\\demo\\go"
	fmt.Println(s2)

	// 转义双引号
	s3 := "hello\"world\""
	fmt.Println(s3)

	// 转义单引号
	s4 := "hello'world'"
	fmt.Println(s4)

	// 常用方法

	// len(s) bytes 字节数

	fmt.Println("bytes", len(str))         // 获取字节数
	fmt.Println("runes", len([]rune(str))) // 获取字符数（rune 数量）

	// 拼接字符串

	// + 或者 fmt.Sprintf

	fmt.Println("拼接", str+s)
	fmt.Println("拼接", fmt.Sprintf("%s %s", str, s)) // %s 格式化字符串

	// strings

	// Split 分割
	s5 := "hello-world-go"

	r := strings.Split(s5, "-") // 返回切片

	fmt.Println(r)

	// Join 拼接

	result := strings.Join(r, "+") // 将切片用分隔符拼接回字符串
	fmt.Println(result)

	// 创建切片

	arr := []string{"php", "java", "python"}
	r1 := strings.Join(arr, "\\")
	fmt.Println(r1)
	fmt.Println(arr)
}
