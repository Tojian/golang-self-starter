package main

import (
	"fmt"
	"time"
)

/**
在Go语言中，每一个并发的执行单元叫作一个goroutine。设想这里的一个程序有两个函数，一个函数做计算，另一个输出结果，假设两个函数没有相互之间的调用关系。
一个线性的程序会先调用其中的一个函数，然后再调用另一个。
如果程序中包含多个goroutine，对两个函数的调用则可能发生在同一时刻。马上就会看到这样的一个程序。
*/

func main() {
	defer testP()
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func testP(){
	fmt.Printf("test test");
}
func spinner(delay time.Duration) {
	for {
		// range 在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
