package main

// import (
// 	"fmt"
// 	// 查看字节大小
// 	"unsafe"
// )

// func main() {
// 	//	数据类型
// 	// 基本数据类型 和 复合数据类型

// 	// 基本数据类型
// 	// 整形 int  浮点型 布尔型 字符串

// 	// 复合数据类型

// 	// 数组 切片 结构体 函数 map 通道(channel) 接口

// 	var num int8 = 10

// 	fmt.Printf("value %v type %T \n", num, num)

// 	// 查看占用字节

// 	fmt.Println(unsafe.Sizeof(num))

// 	// 整形

// 	// 有符号整形
// 	// int8、-2^7 到 2^7-1 占位 1 个字节
// 	// int16、-2^15 到 2^15-1 占位 2 个字节
// 	// int32、-2^31 到 2^31-1 占位 4 个字节 [-2147483648, 2147483647]
// 	// int64、-2^63 到 2^63-1 占位 8 个字节  [-9223372036854775808, 9223372036854775807]

// 	// 无符号整形
// 	// uint8、0 到 2^8-1 占位 1 个字节
// 	// uint16、0 到 2^16-1 占位 2 个字节
// 	// uint32、0 到 2^32-1 占位 4 个字节
// 	// uint64 0 到 2^64-1 占位 8 个字节

// 	// int 类型转换

// 	var num1 int16 = 120

// 	var num2 int32 = 1401111111

// 	fmt.Println(num1 + int16(num2))

// 	// 注意高位向低位转换的问题

// 	var n1 int32 = 150

// 	fmt.Printf("value %v\n", int8(n1)) // 转换结果会有问题

// 	// 默认类型
// 	// 在 window 系统为 64 位时 int默认类型为 int64 位 在 32 位时默认字节为 32 位
// 	n2 := 12

// 	fmt.Printf("value %v \n", n2) //value %v 默认值输出

// 	fmt.Printf("type %T\n", n2) // %T 输出类型

// 	fmt.Printf("type %f\n", n2) // %f 输出浮点数类型

// 	fmt.Printf("十进制 %d\n ", n2) // %d 十进制输出

// 	fmt.Printf("二进制 %b\n ", n2) // %b 二进制输出  计算过程（除以 2 取余法）

// 	fmt.Printf("八进制 %o\n ", n2) // %o 八进制输出

// 	fmt.Printf("十六进制 %x\n ", n2) // %o 十六进制输出
// }
