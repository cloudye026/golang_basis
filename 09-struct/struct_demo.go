package main

import (
	"encoding/json"
	"fmt"
)

// ==================== 1. 匿名嵌入（组合） ====================
// 匿名嵌入可以直接访问嵌入结构体的字段，类似"继承"

type Address struct {
	Province string
	City     string
}

type Contact struct {
	Phone string
	Email string
}

// Employee 使用匿名嵌入
type Employee struct {
	Name    string
	Age     int
	Address // 匿名嵌入 - 可直接访问 Province, City
	Contact // 匿名嵌入 - 可直接访问 Phone, Email
}

// ==================== 2 & 3. 结构体方法 + 接收者类型 ====================

// 值接收者：方法内修改不影响原结构体
func (e Employee) GetInfo() string {
	return fmt.Sprintf("%s, %d岁, 来自%s%s", e.Name, e.Age, e.Province, e.City)
}

// 指针接收者：方法内修改会影响原结构体
func (e *Employee) SetAge(newAge int) {
	e.Age = newAge // 这会修改原始结构体
}

// 值接收者尝试修改（不会生效）
func (e Employee) SetAgeFail(newAge int) {
	e.Age = newAge // 这只修改副本，原始结构体不变
}

// ==================== 4. 结构体标签（JSON） ====================

type User struct {
	ID        int    `json:"id"`              // 指定 JSON 字段名
	Username  string `json:"username"`        // 小写
	Password  string `json:"-"`               // 忽略此字段，不序列化
	Email     string `json:"email,omitempty"` // 如果为空则省略
	CreatedAt string `json:"created_at"`      // 下划线命名
}

// TODO(human): 实现一个方法，让 User 验证邮箱格式是否包含 "@"
// func (u *User) ValidateEmail() bool {
//     // 你的代码
// }

func structDemo() {
	fmt.Println("===== 1. 匿名嵌入演示 =====")
	emp := Employee{
		Name: "张三",
		Age:  28,
		Address: Address{
			Province: "广东省",
			City:     "深圳市",
		},
		Contact: Contact{
			Phone: "13800138000",
			Email: "zhangsan@example.com",
		},
	}

	// 匿名嵌入：可以直接访问嵌入结构体的字段
	fmt.Println("直接访问:", emp.Province, emp.City)   // 不需要 emp.Address.Province
	fmt.Println("完整路径:", emp.Address.Province)     // 也可以这样访问
	fmt.Println("联系方式:", emp.Phone, emp.Email)

	fmt.Println("\n===== 2 & 3. 方法和接收者演示 =====")
	fmt.Println("初始信息:", emp.GetInfo())

	// 指针接收者 - 修改生效
	emp.SetAge(30)
	fmt.Println("SetAge(30) 后:", emp.Age) // 输出: 30

	// 值接收者 - 修改不生效
	emp.SetAgeFail(25)
	fmt.Println("SetAgeFail(25) 后:", emp.Age) // 仍然是: 30

	fmt.Println("\n===== 4. 结构体标签（JSON）演示 =====")
	user := User{
		ID:        1,
		Username:  "admin",
		Password:  "secret123", // 这个不会出现在 JSON 中
		Email:     "",          // 空值，omitempty 会省略
		CreatedAt: "2024-01-01",
	}

	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("序列化结果:")
	fmt.Println(string(jsonData))

	// 反序列化
	jsonStr := `{"id": 2, "username": "test", "email": "test@go.com"}`
	var user2 User
	json.Unmarshal([]byte(jsonStr), &user2)
	fmt.Println("\n反序列化结果:", user2)
}
