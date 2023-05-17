package main

import (
	"fmt"
)

func TestPrintMain() {
	fmt.Print("我是控制台打印，我不换行，可以自己控制换行")
	fmt.Println(`我是控制台打印，我换行`)
	fmt.Printf("我是控制台打印，%s \n", "这是格式化占位符信息,可以自己控制换行")
}

// ----------------------------

type User struct {
	Id int64
}

func TestPlaceholder() {
	user := &User{
		Id: 1,
	}
	// 1.通用占位符
	// 值的默认格式表示
	fmt.Printf("%v\n", user)
	// 类似%v，但输出结构体时会添加字段名
	fmt.Printf("%+v\n", user)
	// 值的Go语法表示
	fmt.Printf("%#v\n", user)
	// 打印值的类型
	fmt.Printf("%T\n", user)
	// 百分号
	fmt.Printf("%%\n")

	// 2.布尔型
	fmt.Printf("%t\n", true)

	// 3.整形
	n := 180
	// 表示为二进制
	fmt.Printf("%b\n", n)
	// 该值对应的unicode码值
	fmt.Printf("%c\n", n)
	// 表示为十进制
	fmt.Printf("%d\n", n)
	// 表示为八进制
	fmt.Printf("%o\n", n)
	// 表示为十六进制，使用a-f
	fmt.Printf("%x\n", n)
	// 表示为十六进制，使用A-F
	fmt.Printf("%X\n", n)
	// 表示为Unicode格式：U+1234，等价于”U+%04X”
	fmt.Printf("%U\n", n)
	a := 96
	fmt.Printf("%q\n", a)
	// 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	fmt.Printf("%q\n", 0x4E2D)

	// 4.浮点数与复数
	f := 18.54
	// 无小数部分、二进制指数的科学计数法，如-123456p-78
	fmt.Printf("%b\n", f)
	// 科学计数法，如-1234.456e+78
	fmt.Printf("%e\n", f)
	// 科学计数法，如-1234.456E+78
	fmt.Printf("%E\n", f)
	// 有小数部分但无指数部分，如123.456
	fmt.Printf("%f\n", f)
	// 等价于%f
	fmt.Printf("%F\n", f)
	// 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	fmt.Printf("%g\n", f)
	// 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	fmt.Printf("%G\n", f)

	// 5.字符串和[]byte
	s := "我是字符串"
	b := []byte{65, 66, 67}
	// 直接输出字符串或者[]byte
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)
	// 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	fmt.Printf("%q\n", s)
	// 每个字节用两字符十六进制数表示（使用a-f
	fmt.Printf("%x\n", s)
	// 每个字节用两字符十六进制数表示（使用A-F）
	fmt.Printf("%X\n", s)

	// 6.指针 占位符 %p 表示为十六进制，并加上前导的0x

	// 7.宽度标识符
	// %10.2
	//宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。
	//精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。
	n1 := 13.14
	// 默认宽度，默认精度
	fmt.Printf("%f\n", n1)
	// 宽度9，默认精度
	fmt.Printf("%10f\n", n1)
	fmt.Printf("%10s\n", "我是字符串")
	// 默认宽度，精度2
	fmt.Printf("%.2f\n", n1)
	// 宽度9，精度2
	fmt.Printf("%10.2f\n", n1)
	// 宽度9，精度0
	fmt.Printf("%10.f\n", n1)

	// 8.其他falg
	// +	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
	//空格	对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格
	//-	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
	//#	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值；
	//0	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
	s1 := "我是字符串"
	fmt.Printf("%  d\n", 10)
	fmt.Printf("%s\n", s1)
	fmt.Printf("%10s\n", s1)
	fmt.Printf("%-10s\n", s1)
	fmt.Printf("%10.2f\n", 10.14)
	fmt.Printf("%-10.2f\n", 10.14)
	fmt.Printf("%010s\n", s1)

	// ...
}

func main() {
	TestPrintMain()
	TestPlaceholder()

}
