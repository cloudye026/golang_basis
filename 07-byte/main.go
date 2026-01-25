package main

import (
	"fmt"
)

// 定义结构体关键字 struct

type hobby struct {
	Food string
	Cake string
}
type person struct {
	Name  string
	Age   int
	City  string
	hobby hobby
}

// 创建结构体的多种方式
type Person struct {
	Name string
	Age  int
}

func SetAge(p *Person) {
	p.Age = 100
}

// 匿名字段(嵌入)

type Food struct {
	string
	int
}

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
	// 类型别名 (用 =)
	//	type byte = uint8 // byte 和 uint8 完全相同

	// 类型定义 (不用 =)

	type MyInt int // Myint 是新类型, 不能和 int 互换

	var num MyInt = 1

	var num2 int = 2

	// cannot use num (variable of int type MyInt) as int value in assignment
	// 	num2 = num // 编译报错 需要进行转换

	num2 = int(num)

	fmt.Println(num2)

	// 创建结构体实列
	p := person{
		Name: "小明",
		Age:  18,
		City: "北京",
		// 嵌套结构体
		hobby: hobby{
			Food: "芒果",
			Cake: "零食",
		},
	}

	// 访问字段
	fmt.Println(p.Name)
	fmt.Println(p.hobby.Food)

	// 修改字段

	p.Age = 19
	fmt.Println(p.Age) // 输出: 19

	// 1. 指定字段名

	p1 := Person{
		Name: "小明",
		Age:  18,
	}

	// 2. 按顺序赋值(不推荐,容易出错)

	p2 := Person{"小红", 19}

	// 3. 零值初始化

	var p3 Person // Name="" Age=0

	// 4. 使用 new (返回指针)

	p4 := new(Person) // *Person 类型

	p4.Name = "王五"

	// 5. 取地址

	p5 := &Person{
		Name: "小紫",
		Age:  23,
	}

	SetAge(p5)
	fmt.Println("p5 value ", p5)

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	p6 := Food{"菠萝", 2,}
	fmt.Println(p6.string)
}
