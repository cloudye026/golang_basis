package main

import "fmt"

func main() {
	fmt.Println("===== 1. 基本 if 语句 =====")
	age := 18
	if age >= 18 {
		fmt.Println("成年人")
	}

	fmt.Println("\n===== 2. if-else 语句 =====")
	score := 75
	if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	fmt.Println("\n===== 3. if-else if-else 语句 =====")
	grade := 85
	if grade >= 90 {
		fmt.Println("优秀")
	} else if grade >= 80 {
		fmt.Println("良好")
	} else if grade >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	fmt.Println("\n===== 4. if 语句中的初始化 =====")
	// 变量 num 只在 if 作用域内有效
	if num := 10; num > 5 {
		fmt.Printf("num = %d, 大于 5\n", num)
	}
	// fmt.Println(num) // 错误：num 在这里不可见

	fmt.Println("\n===== 5. 实际应用：错误处理 =====")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("结果:", result)
	}

	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("错误:", err2)
	} else {
		fmt.Println("结果:", result2)
	}

	fmt.Println("\n===== 6. 嵌套 if 语句 =====")
	username := "admin"
	password := "123456"
	
	if username == "admin" {
		if password == "123456" {
			fmt.Println("登录成功")
		} else {
			fmt.Println("密码错误")
		}
	} else {
		fmt.Println("用户名不存在")
	}

	fmt.Println("\n===== 7. if 与逻辑运算符 =====")
	x := 15
	// && 表示"且"
	if x > 10 && x < 20 {
		fmt.Println("x 在 10 到 20 之间")
	}

	// || 表示"或"
	if x < 10 || x > 20 {
		fmt.Println("x 不在 10 到 20 之间")
	} else {
		fmt.Println("x 在 10 到 20 之间")
	}

	// ! 表示"非"
	isValid := false
	if !isValid {
		fmt.Println("无效")
	}
}

// 除法函数，返回结果和错误
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为 0")
	}
	return a / b, nil
}
