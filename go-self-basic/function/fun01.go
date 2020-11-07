package main

import (
	"fmt"
	"math"
)

/**
函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。
func name(parameter-list) (result-list) {
    body
}

如果一组形参或返回值有相同的类型，我们不必为每个形参都写出参数类型。下面2个声明是等价的：
func f(i, j, k int, s, t string)                 {  }
func f(i int, j int, k int,  s string, t string) { }


*/
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main()  {
	fmt.Println(hypot(3,4)) // "5"
}
