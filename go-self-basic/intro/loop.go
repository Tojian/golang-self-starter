package main

import "fmt"

func main() {

	/**
	Go 只有一种循环结构：for 循环。

	初始化语句：在第一次迭代前执行
	条件表达式：在每次迭代前求值
	后置语句：在每次迭代的结尾执行
	*/
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// 初始化语句和后置语句是可选的。
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// 此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// 如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑
	for {
		fmt.Println(sum)
	}
}
