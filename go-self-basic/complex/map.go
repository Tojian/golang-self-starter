package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
哈希表是一种巧妙并且实用的数据结构。它是一个无序的key/value对的集合，其中所有的key都是不同的，
然后通过给定的key可以在常数时间复杂度内检索、更新或删除对应的value。
*/


type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	// 内置的make函数可以创建一个map：
	ages := make(map[string]int) // mapping from strings to ints
	ages["alice"] = 31
	ages["charlie"] = 34

	ages["alice"] = 32
	fmt.Println(ages["alice"])

	// 使用内置的delete函数可以删除元素：
	delete(ages, "alice") // remove element ages["alice"]

	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}


	fmt.Println(m)
}
