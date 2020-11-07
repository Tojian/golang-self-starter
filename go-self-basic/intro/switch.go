package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句
*/
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	/**
	  没有条件的 switch 同 switch true 一样。

	  这种形式能将一长串 if-then-else 写得更加清晰。
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
