package main

import "fmt"

/**
数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。因为数组的长度是固定的，
因此在Go语言中很少直接使用数组

类型 [n]T 表示拥有 n 个 T 类型的值的数组。

表达式

var a [10]int

*/

func main() {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var a01 [2]string
	a01[0] = "Hello"
	a01[1] = "World"
	fmt.Println(a01[0], a01[1])
	fmt.Println(a01)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
