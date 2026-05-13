# go-rainbow 中文说明

`go-rainbow` 是一个使用 Go 重写的 ANSI 终端文本着色库，参考了 Ruby 项目 `rainbow` 的核心能力，并按 Go 项目的组织方式重新拆分职责。

这个项目遵循一个明确原则：

> 用尽可能少的逻辑，完成足够完整且容易维护的功能。

当前版本：

```text
v0.1.4
```

## 安装

```bash
# 安装当前稳定版本
go get github.com/wangmings/go-rainbow@v0.1.4
```

## 功能概览

- 支持链式样式调用
- 支持前景色与背景色
- 支持 ANSI 编号色
- 支持命名色
- 支持完整 X11 颜色名
- 支持 RGB 三元组
- 支持十六进制颜色字符串
- 支持全局启停与终端环境自动判断
- 支持独立 `Painter`
- 支持移除 ANSI 控制符
- 使用 MIT License，pkg.go.dev 可以正常展示文档

## 目录结构

```text
.
├── ansi.go
├── color.go
├── demo
│   └── main.go
├── doc.go
├── global.go
├── LICENSE
├── named_colors.go
├── palette.go
├── presenter.go
├── rainbow_test.go
├── version.go
└── go.mod
```

各文件职责如下：

- `demo/main.go`
  - 项目根目录下的控制台演示入口
  - 负责直接组织 demo 调用与输出

- `doc.go`
  - 提供包级 Go Doc
  - 用于 pkg.go.dev 首页说明

- `global.go`
  - 提供全局入口
  - 管理默认启停状态
  - 提供 `Wrap`、`Enabled`、`Uncolor`、`New`
  - 按 TTY、`TERM=dumb`、`CLICOLOR_FORCE=1` 初始化默认颜色状态

- `presenter.go`
  - 定义 `Painter`
  - 定义链式文本对象 `Text`
  - 提供样式方法，例如 `Red`、`Bright`、`Underline`

- `named_colors.go`
  - 提供 X11 颜色快捷方法
  - 例如 `Midnightblue`、`Tomato`、`Aliceblue`

- `color.go`
  - 负责颜色参数解析
  - 负责 RGB、Hex、命名色到 ANSI 编码的转换

- `palette.go`
  - 保存 ANSI 色表
  - 保存 X11 颜色表

- `ansi.go`
  - 负责 ANSI 转义序列拼接
  - 负责移除颜色控制码

- `rainbow_test.go`
  - 覆盖主要行为
  - 确保链式调用和颜色转换保持稳定

- `version.go`
  - 保存当前项目版本号

- `LICENSE`
  - 使用 MIT 许可证
  - 允许 pkg.go.dev 展示包文档

## 基本用法

```go
// 声明当前文件属于 main 包
package main

// 引入标准输出包和 go-rainbow 核心包
import (
	// 用于把彩色文本打印到控制台
	"fmt"

	// 引入当前项目的 rainbow 包
	rainbow "github.com/wangmings/go-rainbow"
)

// 程序入口
func main() {
	// 将普通文本包装为可链式着色对象，再输出红色加粗文本
	fmt.Println(rainbow.Wrap("hello").Red().Bright().String())
}
```

输出内容会被包装成 ANSI 控制序列，终端显示时会带颜色和样式。

## 链式调用

```go
// 使用命名色作为前景色
rainbow.Wrap("red text").Foreground("red")

// 使用 FG 别名设置前景色
rainbow.Wrap("blue text").FG("blue")

// 使用英式拼写 Colour 设置前景色
rainbow.Wrap("colour alias").Colour("red")

// 使用 ANSI 数字索引设置前景色
rainbow.Wrap("indexed").Foreground(5)

// 使用 RGB 三元组设置前景色
rainbow.Wrap("rgb").Foreground(255, 128, 64)

// 使用十六进制字符串设置前景色
rainbow.Wrap("hex").Foreground("#ff8040")
```

项目采用 Go 常见的导出 API 形式，便于在 `main` 包和其他外部包中直接调用。

## 背景色

```go
// 使用命名色设置背景色
rainbow.Wrap("green bg").Background("green")

// 使用 BG 别名和十六进制颜色设置背景色
rainbow.Wrap("orange bg").BG("#ff8040")

// 使用 RGB 三元组设置背景色，并把文字改成白色
rainbow.Wrap("rgb background").Background(32, 96, 180).White()
```

## 链式样式

```go
// 将黄色、加粗、下划线三个效果按顺序叠加
rainbow.Wrap("important").
	// 设置黄色前景色
	Yellow().
	// 设置高亮样式
	Bright().
	// 设置下划线样式
	Underline()
```

已提供的样式方法包括：

- `Bright`
- `Bold`
- `Faint`
- `Dark`
- `Italic`
- `Underline`
- `Blink`
- `Inverse`
- `Hide`
- `CrossOut`
- `Strike`
- `Reset`

## 内置颜色快捷方法

```go
// 直接调用 ANSI 红色快捷方法
rainbow.Wrap("hello").Red()

// 直接调用 ANSI 绿色快捷方法
rainbow.Wrap("hello").Green()

// 直接调用 ANSI 蓝色快捷方法
rainbow.Wrap("hello").Blue()

// 直接调用 ANSI 青色快捷方法
rainbow.Wrap("hello").Cyan()
```

当前快捷方法包括：

- `Black`
- `Red`
- `Green`
- `Yellow`
- `Blue`
- `Magenta`
- `Cyan`
- `White`

完整 X11 颜色也提供快捷方法，例如：

```go
// 使用午夜蓝 X11 快捷色
rainbow.Wrap("x11").Midnightblue()

// 使用番茄红 X11 快捷色
rainbow.Wrap("x11").Tomato()

// 使用爱丽丝蓝 X11 快捷色
rainbow.Wrap("x11").Aliceblue()
```

## 样式能力

```go
// 输出高亮文本
rainbow.Wrap("bright").Bright()

// 输出加粗文本，Bold 是 Bright 的别名
rainbow.Wrap("bold alias").Bold()

// 输出淡化文本
rainbow.Wrap("faint").Faint()

// 输出暗淡文本，Dark 是 Faint 的别名
rainbow.Wrap("dark alias").Dark()

// 输出斜体文本
rainbow.Wrap("italic").Italic()

// 输出带下划线的文本
rainbow.Wrap("underline").Underline()

// 输出闪烁文本
rainbow.Wrap("blink").Blink()

// 输出前景背景反转文本
rainbow.Wrap("inverse").Inverse()

// 输出隐藏文本
rainbow.Wrap("hide").Hide()

// 输出删除线文本
rainbow.Wrap("cross out").CrossOut()

// 输出删除线文本，Strike 是 CrossOut 的别名
rainbow.Wrap("strike alias").Strike()

// 输出显式 reset 样式
rainbow.Wrap("reset").Reset()
```

## 完整链式组合

```go
// 演示前景色、背景色和多种样式的连续叠加
rainbow.Wrap("full chain").
	// 设置红色前景色
	Foreground("red").
	// 开启高亮
	Bright().
	// 再次通过 Bold 别名叠加高亮
	Bold().
	// 开启淡化
	Faint().
	// 通过 Dark 别名再次叠加淡化
	Dark().
	// 开启斜体
	Italic().
	// 设置十六进制背景色
	Background("#ff8040").
	// 添加下划线
	Underline().
	// 再把前景色改成蓝色
	Color("blue").
	// 开启闪烁
	Blink().
	// 开启反色
	Inverse().
	// 开启隐藏
	Hide().
	// 添加删除线
	CrossOut().
	// 通过 Strike 别名再次添加删除线
	Strike()
```

## 全局启停

```go
// 关闭全局颜色输出
rainbow.Enabled(false)

// 在禁用状态下创建文本包装对象
text := rainbow.Wrap("hello")

// 打印禁用状态下的文本，输出会保持原文
fmt.Println(text.String())

// 禁用状态下即使传入非法颜色，也会直接返回原文本
fmt.Println(rainbow.Wrap("disabled ignores invalid color").Color("not-a-real-color").Background("still-not-real").String())
```

关闭后，文本会保持原样，不再插入 ANSI 控制序列。
此时颜色参数不会被解析，因此即使传入无效颜色，也会像原项目的 disabled 模式一样直接返回原文本。

恢复启用：

```go
// 重新开启全局颜色输出
rainbow.Enabled(true)
```

## 独立 Painter

如果你不希望所有调用共享同一套启停状态，可以创建独立实例：

```go
// 基于当前全局启停状态创建新的 Painter
p := rainbow.New()

// 使用独立 Painter 包装文本，并输出 X11 颜色与加粗样式
fmt.Println(p.Wrap("hello").Green().Bright().String())
```

关闭某个实例：

```go
// 创建一个永远不输出颜色的独立 Painter
p := rainbow.Painter{Enabled: false}

// 该 Painter 即使调用颜色和样式方法，也会保持原始文本
fmt.Println(p.Wrap("disabled custom painter").Tomato().Underline().String())
```

`rainbow.New()` 会继承当前全局 `Enabled` 状态，之后实例和全局状态互不影响。

## 去除 ANSI 控制符

```go
// 生成带颜色和加粗控制码的原始字符串
raw := "\x1b[31mhello\x1b[0m"

// 删除 ANSI SGR 控制码，得到纯文本
plain := rainbow.Uncolor(raw)
```

`plain` 会得到不带终端样式控制码的普通字符串。

## 设计说明

这个 Go 版本没有照搬 Ruby 的动态能力，而是做了更适合 Go 的取舍：

- 通过 `Text` 结构体实现链式调用
- 通过普通函数和显式类型保持行为清晰
- 通过拆分文件让职责单一
- 通过少量内部辅助函数复用核心逻辑
- 通过测试固定关键行为，避免后续重构失真

Ruby 原项目里的 `String` mixin 与 refinement 属于 Ruby 语言特性，Go 不能给内建 `string` 动态追加方法。
Go 版本用 `rainbow.Wrap(...)` 作为等价入口，其他可迁移的颜色、样式、实例状态与环境行为均已补齐。

## 当前状态

项目当前已经具备一个可工作的最小核心版本，并通过测试：

```bash
go test ./...
```

pkg.go.dev 文档地址：

```text
https://pkg.go.dev/github.com/wangmings/go-rainbow@v0.1.4
```

## 命令行 Demo

在项目根目录执行：

```bash
go run ./demo
```

可以直接在终端看到以下能力：

- 基础颜色
- ANSI 数字色
- `FG`、`BG`
- `Color`、`Colour`
- 完整 X11 快捷色
- 全部样式能力
- 完整链式组合
- 全局启停
- 独立 `Painter`
- Hex 前景色
- RGB 背景色
- `Uncolor` 去除控制符

如果当前终端环境被自动判断为非彩色输出，可以使用下面的命令强制查看完整彩色演示：

```bash
# 强制启用颜色输出，再运行项目入口 demo
CLICOLOR_FORCE=1 go run ./demo
```

## 后续可扩展方向

- 提供返回 `error` 的安全 API
- 增加英文 README
