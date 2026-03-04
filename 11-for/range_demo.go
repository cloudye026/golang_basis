package main

import "fmt"

func main() {
	fmt.Println("===== range 关键字详解 =====\n")
	
	// range 是 Go 的内置关键字，用于遍历（迭代）各种数据结构
	// 它会返回两个值：索引/键 和 值
	
	fmt.Println("1. range 遍历切片 - 返回：索引 + 值")
	fruits := []string{"苹果", "香蕉", "橙子"}
	for index, value := range fruits {
		fmt.Printf("  索引 %d: %s\n", index, value)
	}
	
	fmt.Println("\n2. range 遍历数组 - 返回：索引 + 值")
	numbers := [3]int{10, 20, 30}
	for i, num := range numbers {
		fmt.Printf("  numbers[%d] = %d\n", i, num)
	}
	
	fmt.Println("\n3. range 遍历 map - 返回：键 + 值")
	ages := map[string]int{
		"张三": 25,
		"李四": 30,
	}
	for name, age := range ages {
		fmt.Printf("  %s 的年龄是 %d\n", name, age)
	}
	
	fmt.Println("\n4. range 遍历字符串 - 返回：字节索引 + Unicode 码点")
	text := "Go语言"
	for index, char := range text {
		fmt.Printf("  索引 %d: 字符 '%c' (Unicode: %U)\n", index, char, char)
	}
	
	fmt.Println("\n5. range 遍历 channel - 返回：值（只有一个返回值）")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // 必须关闭，否则会死锁
	
	for value := range ch {
		fmt.Printf("  从 channel 接收: %d\n", value)
	}
	
	fmt.Println("\n===== range 的使用技巧 =====\n")
	
	fmt.Println("技巧 1: 只要索引，忽略值")
	for i := range fruits {
		fmt.Printf("  索引: %d\n", i)
	}
	
	fmt.Println("\n技巧 2: 只要值，忽略索引（用 _ 占位）")
	for _, fruit := range fruits {
		fmt.Printf("  水果: %s\n", fruit)
	}
	
	fmt.Println("\n技巧 3: 两个都不要（只是执行 N 次）")
	count := 0
	for range fruits {
		count++
	}
	fmt.Printf("  切片长度: %d\n", count)
	
	fmt.Println("\n===== range 的底层原理 =====\n")
	
	// range 会复制被遍历的值
	fmt.Println("示例：修改 range 返回的值不会影响原数据")
	nums := []int{1, 2, 3}
	for _, num := range nums {
		num = num * 10 // 修改的是副本
	}
	fmt.Printf("  原始数据: %v (没有变化)\n", nums)
	
	fmt.Println("\n如果要修改原数据，使用索引:")
	for i := range nums {
		nums[i] = nums[i] * 10 // 通过索引修改
	}
	fmt.Printf("  修改后: %v\n", nums)
	
	fmt.Println("\n===== range vs 传统 for 循环 =====\n")
	
	data := []int{10, 20, 30, 40, 50}
	
	fmt.Println("传统 for 循环:")
	for i := 0; i < len(data); i++ {
		fmt.Printf("  data[%d] = %d\n", i, data[i])
	}
	
	fmt.Println("\nrange 循环（更简洁）:")
	for i, v := range data {
		fmt.Printf("  data[%d] = %d\n", i, v)
	}
}
