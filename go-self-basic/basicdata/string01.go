package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(len(s))
	fmt.Print(s[0], s[7])
	/**
	  panic: runtime error: index out of range [11] with length 11
	*/
	// c := s[len(s)]
	// fmt.Print(c)

	fmt.Println(s[0:5]) // "hello"
	// 不管i还是j都可能被忽略，当它们被忽略时将采用0作为开始位置，采用len(s)作为结束的位置。
	fmt.Println(s[:5]) // "hello"
	fmt.Println(s[7:]) // "world"
	fmt.Println(s[:])  // "hello, world"

	//
	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	// 追加
	s01 := "left foot"
	t := s01
	s01 += ", right foot"
	fmt.Print(s01) // "left foot, right foot"
	fmt.Print(t) // "left foot"

	// 因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的：
}
