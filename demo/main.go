package main

import (
	"fmt"

	rainbow "github.com/wangmings/go-rainbow"
)

func main() {
	// 输出前景色输入方式分组标题
	section("Color input")
	// 使用 ANSI 基础色快捷方法输出红色文本
	line(rainbow.Wrap("ANSI shortcut").Red())
	// 使用 ANSI 数字索引设置前景色
	line(rainbow.Wrap("ANSI number foreground").Foreground(5))
	// 使用 FG 别名设置青色前景色
	line(rainbow.Wrap("ANSI alias FG").FG("cyan"))
	// 直接通过 Color 设置黄色前景色
	line(rainbow.Wrap("direct Color").Color("yellow"))
	// 使用英式拼写 Colour 设置番茄红前景色
	line(rainbow.Wrap("Colour alias").Colour("tomato"))
	// 使用十六进制字符串设置前景色
	line(rainbow.Wrap("hex foreground").Foreground("#ff8040"))
	// 使用 RGB 三元组设置前景色
	line(rainbow.Wrap("RGB foreground").Foreground(255, 128, 64))
	// 使用 X11 快捷方法设置午夜蓝前景色
	line(rainbow.Wrap("named X11 color").Midnightblue())

	// 输出背景色输入方式分组标题
	section("Background input")
	// 使用命名色设置绿色背景
	line(rainbow.Wrap("named background").Background("green"))
	// 使用 BG 别名设置蓝色背景，并把文字改成白色
	line(rainbow.Wrap("ANSI alias BG").BG(4).White())
	// 使用 RGB 三元组设置背景色，并把文字改成白色
	line(rainbow.Wrap("RGB background").Background(32, 96, 180).White())
	// 使用十六进制字符串设置背景色，并把文字改成白色
	line(rainbow.Wrap("hex background").Background("#663399").White())

	// 输出文本样式分组标题
	section("Text styles")
	// 输出高亮文本
	line(rainbow.Wrap("bright").Bright())
	// 输出加粗文本，Bold 是 Bright 的别名
	line(rainbow.Wrap("bold alias").Bold())
	// 输出淡化文本
	line(rainbow.Wrap("faint").Faint())
	// 输出暗淡文本，Dark 是 Faint 的别名
	line(rainbow.Wrap("dark alias").Dark())
	// 输出斜体文本
	line(rainbow.Wrap("italic").Italic())
	// 输出带下划线文本
	line(rainbow.Wrap("underline").Underline())
	// 输出闪烁文本
	line(rainbow.Wrap("blink").Blink())
	// 输出前景色与背景色反转文本
	line(rainbow.Wrap("inverse").Inverse())
	// 输出隐藏文本
	line(rainbow.Wrap("hide").Hide())
	// 输出删除线文本
	line(rainbow.Wrap("cross out").CrossOut())
	// 输出删除线文本，Strike 是 CrossOut 的别名
	line(rainbow.Wrap("strike alias").Strike())
	// 输出显式 reset 文本
	line(rainbow.Wrap("reset").Reset())
	// 输出完整链式组合示例
	line(rainbow.Wrap("full chain").Foreground("red").Bright().Bold().Faint().Dark().Italic().Background("#ff8040").Underline().Color("blue").Blink().Inverse().Hide().CrossOut().Strike())

	// 输出全局启停分组标题
	section("Enable switch")
	// 关闭全局颜色输出
	rainbow.Enabled(false)
	// 在全局禁用状态下输出文本，结果会保持原文
	line(rainbow.Wrap("disabled global painter").Red().Bright())
	// 在禁用状态下传入非法颜色，也不会报错
	line(rainbow.Wrap("disabled ignores invalid color").Color("not-a-real-color").Background("still-not-real"))
	// 重新开启全局颜色输出
	rainbow.Enabled(true)
	// 重新开启后再次输出绿色高亮文本
	line(rainbow.Wrap("enabled again").Green().Bright())

	// 输出独立 Painter 分组标题
	section("Independent painter")
	// 创建继承当前全局状态的独立 Painter
	custom := rainbow.New()
	// 使用独立 Painter 输出爱丽丝蓝加粗文本
	line(custom.Wrap("new painter inherits global enabled state").Aliceblue().Bold())
	// 创建一个始终禁用颜色输出的独立 Painter
	silent := rainbow.Painter{Enabled: false}
	// 使用禁用的独立 Painter 输出文本，结果会保持原文
	line(silent.Wrap("disabled custom painter").Tomato().Underline())

	// 输出去色能力分组标题
	section("Uncolor")
	// 生成带 ANSI 控制码的绿色加粗文本
	raw := rainbow.Wrap("plain text after uncolor").Green().Bold().String()
	// 打印保留 ANSI 控制码的原始字符串
	fmt.Println("styled :", raw)
	// 打印移除 ANSI 控制码后的纯文本
	fmt.Println("plain  :", rainbow.Uncolor(raw))
}

func section(title string) {
	// 打印演示区块标题
	fmt.Printf("\n[%s]\n", title)
}

func line(text rainbow.Text) {
	// 将 Text 转成字符串并输出到控制台
	fmt.Println(text.String())
}
