# go-study

个人 Go 语言学习仓库，包含基础语法练习、数据结构与算法实现、框架学习（Gin / gee）、RPC 示例、LeetCode 题解，以及一些自用小工具。

> 说明：这是学习/实验性质的代码集合，目录之间相互独立，部分示例可能依赖特定环境（如 SQLite/MySQL、证书文件等）。

## 环境要求

- Go：见 `go.mod`（当前为 Go **1.22**）
- OS：macOS / Windows / Linux 均可（部分示例对平台有差异）

## 快速开始

克隆后在仓库根目录执行：

```bash
go test ./...
```

多数示例是 `main` 包，可直接运行某个目录下的代码，例如：

```bash
go run ./basic_syntax/basic/main
```

> 如果某个示例涉及外部依赖（数据库、网络端口等），请先阅读该目录下的 README 或源码注释。

## 项目结构

下面是核心目录的用途概览（按“学习模块”划分）：

- `basic_syntax/`：Go 基础语法练习（数组、切片、Map、函数、反射、并发、网络等）
- `algorithm/`：算法实现（如 Levenshtein、MinHash）
- `data_structure/`：数据结构实现（LRU、排序、树等）
- `design_pattern/`：设计模式（单例等）
- `gee_framework/`：跟随 gee Web 框架实现的学习代码（含测试与示例）
- `gin_framework/`：Gin 框架的使用示例（中间件、路由、模板、上传等）
- `golang_rpc/`：使用 Go 原生库实现的简单 RPC 示例
- `leetcode_problem/`：LeetCode 题目练习与解法
- `simple_tutorial/`：学习《Go语言简明教程》过程中的代码
- `the-way-to-go/`：学习《Go 入门指南》（learnku 译文）时的笔记与练习
- `calculate_delivery_fee/`：笔试题实现（包含 SQLite 相关代码与测试）
- `tool/`：自用工具（例如语雀 Markdown 清洗/格式化）

## 工具：语雀 Markdown 清洗（`tool/lark_tool.go`）

用于处理语雀导出的 Markdown：

- 将 `<br />` 替换为更标准的换行
- 清理链接中的 `#...` 片段（例如 `https://xxx#xxx`）
- 批量处理指定目录下的同后缀文件（默认 `.md`）

### 使用方式

`lark_tool` 支持通过参数指定扫描目录与行为：

- `-dir`：要扫描的目录（默认：优先使用 `~/Downloads`，若不存在则使用当前工作目录）
- `-ext`：要处理的文件后缀（默认：`.md`）
- `-recursive`：是否递归扫描子目录（默认：macOS 不递归；Windows 默认递归；也可显式传入覆盖）

示例：

```bash
# 直接运行源码
go run ./tool/lark_tool.go -dir "$HOME/Downloads" -ext .md

# 编译后运行（推荐）
go build -o /tmp/lark_tool ./tool/lark_tool.go
/tmp/lark_tool -dir "$HOME/Downloads" -recursive
```

> 为什么“打包后找不到目录”？
> 
> Go 中相对路径基于“进程工作目录（working directory）”，而不是基于源码目录或可执行文件目录。
> 打包后的程序在不同位置运行时，working dir 往往变化，导致相对路径失效。当前版本通过 `-dir` 参数与默认目录策略规避了这个问题，并会打印 working dir / executable dir 便于排查。

## 常用命令

```bash
# 运行所有测试
go test ./...

# 查看/整理依赖
go mod tidy

# 静态检查（如本地已安装）
# go vet ./...
```

## 贡献

这是个人学习仓库，欢迎提 Issue/PR：

- 修复明显 bug 或可读性问题
- 补充更小、更完整的示例或测试用例
- 改善文档结构与使用说明

## License

本项目使用根目录 `LICENSE` 中声明的许可证。
