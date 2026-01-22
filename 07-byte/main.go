package main

import (
	"fmt"
)

func main() {
	// 类型别名
	/**
	byte 是 uint8 的别名
	它们是 完全相同 的类型
	byte 只是 uint8 的另一个名字
	*/
	type byte = uint8 // 取值范围：0 到 255

	// 示例

	var data1 uint8 = 65

	// 更清晰
	var data byte = 65

	fmt.Printf("%T %T\n", data, data1)

	// 字符串和字节切片转换

	str := "Hello"

	bytes := []byte(str) // 字符串转切片

	fmt.Println(bytes)

	str2 := string(bytes) // 字节切片转字符串

	fmt.Println(str2)

	// 类型别名 vs 类型定义

	// 区别:
	//  类型别名 (用 =)
	//	type byte = uint8 // byte 和 uint8 完全相同

	// 类型定义 (不用 =)

	type MyInt int // Myint 是新类型, 不能和 int 互换

	var num MyInt = 1

	var num2 int = 2

	// cannot use num (variable of int type MyInt) as int value in assignment
	// 	num2 = num // 编译报错 需要进行转换

	num2 = int(num)

	fmt.Println(num2)

}
