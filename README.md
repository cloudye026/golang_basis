# Go 基础练习（basis）

面向初学与回顾的 Go 语言基础示例集合。每个子目录聚焦一个主题，包含可运行或可阅读的示例代码与注释说明。

## 环境要求

- Go `1.21.5`（见 `go.mod`）或以上
- 可选：Git（便于克隆与版本管理）

## 目录结构

- `01-variable/` 变量与输出：`fmt.Print/Printf/Println`、格式化占位符、结构体打印、变量声明方式等（当前代码主要用于阅读，需取消注释后运行）。
- `02-int/` 整型：有/无符号整型、位宽与范围、`unsafe.Sizeof`、进制打印、类型转换注意事项（当前代码主要用于阅读，需取消注释后运行）。
- `03-float/` 浮点型：`float32/float64`、精度与格式化、科学计数法表示、精度丢失与 `shopspring/decimal` 的使用（该示例可直接运行）。

## 快速开始

在任意可运行示例目录执行 `go run`：

```bash
# 进入仓库根目录后
# 运行浮点型示例
go run ./03-float
```

如需运行 `01-variable` 或 `02-int`，先打开对应目录下的 `main.go` / `*.go` 文件，按需取消注释的 `import`、`main()` 和示例代码，然后再执行：

```bash
go run ./01-variable
# 或
go run ./02-int
```

## 常用命令

```bash
# 运行某个示例目录
go run ./03-float

# 构建某个示例目录
go build ./03-float

# 代码格式化与静态检查（可在仓库根目录执行）
go fmt ./...
go vet ./...
```

## 学习要点速览

- 输出与格式化：`%v`、`%+v`、`%#v`、`%d`、`%b`、`%o`、`%x`、`%T`、`%.2f` 等常用占位符。
- 变量声明：`var` 显式类型、类型推导与 `:=` 短变量声明（仅函数内）。
- 整型细节：不同位宽的取值范围、默认位宽、不同进制打印与显示。
- 类型转换：高位向低位转换可能溢出；`float` → `int` 会截断小数，不建议随意转换。
- 浮点精度：二进制浮点的舍入误差；金额/高精度计算可使用 `github.com/shopspring/decimal`。

## 依赖说明

- `github.com/shopspring/decimal`：在 `03-float` 中用于演示十进制高精度计算（加/减/乘/除）。

### 扩展与新增示例

建议以序号创建新主题目录，例如 `04-bool/`、`05-string/`、`06-slice/` 等，并在其中添加 `main.go`：

```go
package main

import "fmt"

func main() {
    fmt.Println("demo")
}
```

然后使用 `go run ./04-bool` 运行。

### 参考

- Go 官方文档与语言规范：https://go.dev/doc/
- `shopspring/decimal` 项目主页：https://github.com/shopspring/decimal

