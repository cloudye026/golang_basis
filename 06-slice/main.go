package main

import "fmt"

func main() {
	// ========== 切片的定义 ==========

	// 方式1: 使用 make 函数创建切片
	// make([]类型, 长度, 容量)
	slice1 := make([]int, 5)       // 长度和容量都是5
	slice2 := make([]int, 3, 10)   // 长度为3，容量为10
	fmt.Println("slice1:", slice1) // [0 0 0 0 0]
	fmt.Println("slice2:", slice2) // [0 0 0]
	fmt.Printf("slice2 长度: %d, 容量: %d\n", len(slice2), cap(slice2))

	// 方式2: 直接声明并初始化
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice3:", slice3)

	// 方式3: 从数组创建切片
	arr := [5]int{10, 20, 30, 40, 50}
	slice4 := arr[1:4] // 从索引1到3（不包含4）
	fmt.Println("slice4:", slice4) // [20 30 40]

	// 方式4: 声明一个空切片（nil切片）
	var slice5 []int
	fmt.Println("slice5 是否为nil:", slice5 == nil) // true

	// ========== 切片的操作 ==========

	// 追加元素
	slice3 = append(slice3, 6, 7, 8)
	fmt.Println("追加后的 slice3:", slice3)

	// 切片的切片操作
	subSlice := slice3[2:5]
	fmt.Println("subSlice:", subSlice)

	// 遍历切片
	fmt.Println("遍历切片:")
	for index, value := range slice3 {
		fmt.Printf("索引: %d, 值: %d\n", index, value)
	}

	// ========== 切片与数组的区别 ==========
	// 1. 数组长度固定，切片长度可变
	// 2. 数组是值类型，切片是引用类型
	// 3. 切片底层引用数组

	// 切片是引用类型的证明
	original := []int{1, 2, 3}
	copied := original     // 这是引用复制，不是值复制
	copied[0] = 999
	fmt.Println("original:", original) // [999 2 3] - 原切片也被修改了
	fmt.Println("copied:", copied)     // [999 2 3]
}
