package main

// import "fmt"

// // 定义结构体类型
// type Object struct {
// 	A int
// 	B int
// 	C int
// }

// func main() {
// 	// fmt.Print("111")
// 	// fmt.Print("111")
// 	// fmt.Print("111")
// 	var a string = "123"
// 	// Print 直接输出
// 	// 而 Print 版本则仅在操作数两侧都不是字符串时才添加空格
// 	fmt.Print(a)

// 	// Printf 使用不同的参数类型来决定这些属性
// 	// 通用格式 %v 该格式可以打印任何值
// 	// %+v 会用其名称注释结构的字段
// 	// 替代格式 %#v 会以完整的 Go 语法打印该值
// 	fmt.Printf("%v", "222")
// 	fmt.Printf("%v", "222")

// 	// Println 换行输出
// 	// Println 还会在参数之间插入空格，并在输出后附加换行符
// 	// fmt.Println("333")
// 	// fmt.Println("333")
// 	// fmt.Println("333")

// 	// 创建结构体实例
// 	var obj Object = Object{A: 1, B: 2, C: 3}

// 	// 打印结构体
// 	fmt.Printf("\n%v\n", obj) // 普通格式
// 	fmt.Printf("%+v\n", obj)  // 带字段名
// 	fmt.Printf("%#v\n", obj)  // 完整 Go 语法

// 	// 变量类型推导（三种方式）

// 	// 方式 1：var + 显式类型
// 	var x int = 100

// 	// 方式 2：var + 类型推导
// 	var y = 200

// 	// 方式 3：:= 短声明（最常用） 只能在函数内部使用ß
// 	z := 300

// 	value := x

// 	fmt.Printf("\nx=%d, y=%d, z=%d\n", x, y, z)
// 	fmt.Printf("%#v\n", value)

// 	// 变量声明

// 	n := "123"
// 	fmt.Println(n)
// }
