package main

import "fmt"

/**
一、变量
主要利用关键字 go
格式：
 var 变量名字 类型 = 表达式
 var s string
 var i, j, k int
 var b, f, s = true, 2.3, "four"


或者直接用这种	b := "123" 进行申明和初始化

var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。

就像在这个例子中看到的一样，var 语句可以出现在包或函数级别。


变量声明可以包含初始值，每个变量对应一个。

如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。

在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用
*/

var i, j int = 1, 2

func main() {
	fmt.Print("hello")

	var i int
	i = 0

	b := "123"
	var s = "hello"
	var s01 string = "hello01"
	fmt.Print(s)
	fmt.Print(s01)
	fmt.Print(i)
	fmt.Print(b)

	// 变量的初始化
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
