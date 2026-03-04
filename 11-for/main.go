package main

import "fmt"

func main() {
	fmt.Println("===== 1. 基本 for 循环（类似 C 语言）=====")
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println("\n===== 2. for 作为 while 循环 =====")
	// Go 没有 while 关键字，用 for 代替
	count := 0
	for count < 3 {
		fmt.Printf("count = %d\n", count)
		count++
	}

	fmt.Println("\n===== 3. 无限循环 =====")
	n := 0
	for {
		n++
		fmt.Printf("n = %d\n", n)
		if n >= 3 {
			break // 使用 break 跳出循环
		}
	}

	fmt.Println("\n===== 4. range 遍历数组/切片 =====")
	nums := []int{10, 20, 30, 40, 50}
	
	// 同时获取索引和值
	for index, value := range nums {
		fmt.Printf("索引: %d, 值: %d\n", index, value)
	}

	// 只要值，忽略索引
	fmt.Println("\n只要值:")
	for _, value := range nums {
		fmt.Printf("值: %d\n", value)
	}

	// 只要索引
	fmt.Println("\n只要索引:")
	for index := range nums {
		fmt.Printf("索引: %d\n", index)
	}

	fmt.Println("\n===== 5. range 遍历字符串 =====")
	str := "Hello世界"
	for index, char := range str {
		fmt.Printf("索引: %d, 字符: %c, Unicode: %U\n", index, char, char)
	}

	fmt.Println("\n===== 6. range 遍历 map =====")
	scores := map[string]int{
		"张三": 85,
		"李四": 92,
		"王五": 78,
	}
	
	for name, score := range scores {
		fmt.Printf("%s: %d 分\n", name, score)
	}

	fmt.Println("\n===== 7. continue 跳过本次循环 =====")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // 跳过 i=3
		}
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println("\n===== 8. break 跳出循环 =====")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break // 在 i=3 时跳出
		}
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println("\n===== 9. 嵌套循环 =====")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}

	fmt.Println("\n===== 10. 标签跳出多层循环 =====")
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d, %d) ", i, j)
			if i == 2 && j == 2 {
				break outer // 跳出外层循环
			}
		}
		fmt.Println()
	}
	fmt.Println("\n已跳出外层循环")

	fmt.Println("\n===== 11. 实际应用：查找元素 =====")
	numbers := []int{5, 12, 8, 20, 15}
	target := 20
	found := false
	
	for _, num := range numbers {
		if num == target {
			found = true
			break
		}
	}
	
	if found {
		fmt.Printf("找到了 %d\n", target)
	} else {
		fmt.Printf("没找到 %d\n", target)
	}

	fmt.Println("\n===== 12. 实际应用：累加求和 =====")
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("1 到 10 的和: %d\n", sum)
}
