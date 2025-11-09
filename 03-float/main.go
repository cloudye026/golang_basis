package main

import (
	"fmt"
	"unsafe"

	"github.com/shopspring/decimal"
)

func main() {
	// 数据类型
	// 浮点型
	// 两种 float32 和 float64
	// 定义 float 类型

	var a float32 = 3.121592653

	fmt.Println(a)
	// 浮点数 32 位占 4 个字节
	fmt.Println(unsafe.Sizeof(a))

	// 保留 小数 %.2f 两位 %.4f 四位 四舍五入

	fmt.Printf("%.2f\n%.4f \n", a, a)

	// 64位 系统中 go 语言默认 float 为 64 位
	f64 := 3.121592653

	fmt.Printf("%T\n", f64)

	num := 120

	fmt.Printf("%T\n", num)

	// 使用科学计数法表示浮点数类型
	f3 := 3.14e2 //表示 3.14*10的2次方

	fmt.Printf("%T\n %f\n", f3, f3)

	f4 := 3.14e-2 //表示 3.14除以10的2次方

	fmt.Printf("%T\n %f\n", f4, f4)

	// float 精度丢失的问题
	f5 := 1199.6 // 119959.99999999999
	fmt.Println(f5 * 100)

	n1 := 8.2
	n2 := 3.8
	fmt.Println(n1 - n2) // 4.3999999999999995

	// Add 加法 Sub 减法 Mul 乘法 Div 除法
	n3 := decimal.NewFromFloat(n1).Sub(decimal.NewFromFloat(n2))

	fmt.Println(n3)


	// 类型转换 int 类型 转 float 类型 

	a1 := 10
	b1 := float64(a1)

	fmt.Printf("%v\n%T\n%v\n%T\n", a1, a1, b1, b1)

	// float 转 int 类型 (不建议转换)

	f6 := 3.14
  i1 := int(f6)

	fmt.Printf("%v\n%T\n%v\n%T", f6, f6, i1,i1)
}
