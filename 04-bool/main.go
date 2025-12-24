package main

import "fmt"

func main() {
	// 布尔类型
	// go 语言中 bool 只有 true 和 false 两个值
	// go 语言中不允许将整形强制转换为 bool 类型
	// bool 类型无法参与数值运算, 也无法与其他类型进行转换
	// bool 类型的变量默认值为 false

	var b bool = true
	fmt.Println(b)

	var b1 bool
	fmt.Println(b1)

	// string 类型默认值为空

	var s string
	fmt.Println(s)

	// int类型默认值为0

	var i int
	fmt.Println(i)

	//  float类型默认值为0

	var f float64
	fmt.Println(f)

	// if 语句
	// bool 无法参与与其他类型的强制转换 也无法参与数值运算

	a := true

	if a {
		println("auture")
	} else {
		println("not auture")
	}
}
