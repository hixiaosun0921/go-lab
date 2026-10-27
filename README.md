# project01 — Go 基础练习

每个练习一个独立目录、独立 `package main`，可以单独编译运行，互不干扰。

```bash
go run ./examples/hello          # 运行单个示例
go build ./...                   # 编译全部示例
```

## 目录

| 目录 | 内容 |
| --- | --- |
| `examples/variables` | 变量定义方式 |
| `examples/arithmetic` | 算术运算 |
| `examples/bit_shift` | 位运算与移位 |
| `examples/escape` | 转义字符 |
| `examples/hello` | Hello World |
| `examples/hello_golang` | Hello Golang |
| `examples/ascii_art` | 字符画 |
| `examples/if_age` | if 判断（年龄） |
| `examples/if_leap_year` | if 判断（闰年） |
| `examples/switch_month` | switch（季节） |
| `examples/for_loop` | for 循环 |
| `examples/for_pyramid` | for 嵌套（打印三角形） |
| `examples/for_exercise` | for 循环练习 |
| `examples/input` | 键盘输入 |
| `examples/pointer` | 指针基础 |
| `examples/swap_pointers` | 指针交换变量 |
| `examples/exchange_values` | 多返回值交换 |
| `examples/variadic_sum` | 可变参数求和 |
| `examples/recursion` | 递归 |
| `examples/recursion_peach` | 递归（桃子问题） |
| `examples/fibonacci` | 递归（斐波那契） |
| `examples/closure` | 闭包 |
| `examples/closure_suffix` | 闭包（后缀补全） |
| `examples/anonymous_func` | 匿名函数 |
| `examples/defer_demo` | defer |
| `examples/break_random` | break（随机数） |
| `examples/string_demo` | 字符串与 model 包 |
| `examples/http_server` | net/http 简单接口 |

`model/` 是可复用包，示例通过 `project01/model` 导入。
