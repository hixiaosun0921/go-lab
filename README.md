# project01 — Go 基础练习

Go 语言基础语法的分主题练习集。共 32 个示例 + 1 个公共包。

每个练习是一个独立目录、独立 `package main`，可以单独编译运行，互不干扰。

## 环境要求

- Go 1.26 或更高版本（用 `go version` 查看）

## 快速开始

```bash
go build ./...                   # 编译全部示例
go run ./examples/hello          # 运行单个示例
go run ./examples/input          # 需要键盘输入的示例（在终端里交互）
```

## 目录

### 入门与基础语法

| 目录 | 内容 |
| --- | --- |
| `examples/hello` | Hello World |
| `examples/hello_golang` | 打印 Hello Golang |
| `examples/ascii_art` | 字符画输出 |
| `examples/variables` | 变量的多种定义方式 |
| `examples/arithmetic` | 算术运算 |
| `examples/bit_shift` | 位运算与移位 |
| `examples/escape` | 转义字符 |
| `examples/zuoyongyv` | 变量作用域（全局变量与局部遮蔽） |

### 流程控制

| 目录 | 内容 |
| --- | --- |
| `examples/if_age` | if 判断（按年龄分组） |
| `examples/if_leap_year` | if 判断（闰年） |
| `examples/switch_month` | switch 判断季节 |
| `examples/for_loop` | for 循环 |
| `examples/for_pyramid` | for 嵌套（打印三角形） |
| `examples/for_exercise` | for 循环综合练习 |
| `examples/break_random` | break 与随机数 |

### 函数与递归

| 目录 | 内容 |
| --- | --- |
| `examples/exchange_values` | 多返回值交换 |
| `examples/variadic_sum` | 可变参数求和 |
| `examples/anonymous_func` | 匿名函数 |
| `examples/closure` | 闭包 |
| `examples/closure_suffix` | 闭包（文件名后缀补全） |
| `examples/defer_demo` | defer 执行顺序 |
| `examples/recursion` | 递归入门 |
| `examples/recursion_peach` | 递归（猴子吃桃问题） |
| `examples/fibonacci` | 递归（斐波那契数列） |

### 指针与输入

| 目录 | 内容 |
| --- | --- |
| `examples/pointer` | 指针基础 |
| `examples/swap_pointers` | 用指针交换两个变量的值 |
| `examples/input` | 键盘输入（`fmt.Scanln`） |

### 字符串与时间

| 目录 | 内容 |
| --- | --- |
| `examples/string_demo` | 字符串不可变性，以及导入 `model` 包 |
| `examples/str_demo` | 字符串常用操作：`len`、`rune`、`strconv`、`strings` |
| `examples/time` | `time.Now()` 与年月日取值 |

### 错误处理

| 目录 | 内容 |
| --- | --- |
| `examples/error1` | `defer` + `recover` 捕获运行时 panic（除零） |
| `examples/error2` | `errors.New` 自定义错误与 `panic` |

## model 包

`model/` 是可复用包，示例如下：

```go
import "project01/model"

fmt.Println(model.Name)   // hhh123
```

调用方是 `examples/string_demo`。

## 约定

- 每个 `examples/*` 目录都是独立的 `package main`，互不依赖，可以单独 `go run`
- 只有 `model/` 是 `package model`，被示例通过模块路径 `project01/model` 导入
- 编译产物由 `.gitignore` 排除（`*.exe`、`*.test`、`*.out` 等）
