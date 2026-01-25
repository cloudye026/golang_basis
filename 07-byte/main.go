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

// 结构体方法

// 定义方法
func (p Person) SayHello() {
	fmt.Println("hello my name is %s\n", p.Name)
}


func (f Food) MyHobby() {
	fmt.Println("i like food is ", f.string)
}

// 修改字段需要用指针接收
// 指针接收者 方法内修改会影响原始结构体

func (p *Person) EditPersonAge(newAge int) {
	 p.Age = newAge
}

// 用户信息
type User struct {
    ID       int
    Username string
    Email    string
    Password string
}

// 商品信息
type Product struct {
    ID    int
    Name  string
    Price float64
    Stock int
}

// 订单信息
type Order struct {
    OrderID   string
    User      User      // 嵌套
    Products  []Product // 切片
    Total     float64
    CreatedAt string
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

	fmt.Println("log p1",p1)
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


	p7 := Person{
		Name: "结构方法体",
	}

	p8 := Food{"芒果",3,}

	p7.SayHello()

	p8.MyHobby()

	// p1 是值类型，但Go会自动转换为指针调用指针方法
	p1.EditPersonAge(20) // 注意调用顺序
	fmt.Println("edit age then value is " , p1.Age)



	
}
