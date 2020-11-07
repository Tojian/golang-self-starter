package main

import (
	"fmt"
	"math/cmplx"
)

/**
基础数据类型

1. 整型

有符号和无符号类型的整数运算。这里有int8、int16、int32和int64四种截然不同大小的有符号整形数类型，
分别对应8、16、32、64bit大小的有符号整形数，与此对应的是uint8、uint16、uint32和uint64四种无符号整形数类型。

int8 、 int16 、 int32 、 int64
uint8、uint16、uint32、uint64

其中有符号整数采用2的补码形式表示，也就是最高bit位用作表示符号位，
一个n-bit的有符号数的值域是从到。无符号整数的所有bit位都用于表示非负数，值域是0到。
例如，int8类型整数的值域是从-128到127，而uint8类型整数的值域是从0到255

2. 浮点数
3. 复数
4. 布尔
5. 字符串
6. 常量

基础类型：
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func basicdata() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func main()  {
	basicdata()
}
