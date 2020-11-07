package main

/**
包和文件
Go语言中的包和其他语言的库或模块的概念类似，目的都是为了支持模块化、封装、单独编译和代码重用。
 */
// CToF converts a Celsius temperature to Fahrenheit.
func CToF01(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC01(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }