package main

import "fmt"

/**
Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。
一个slice类型一般写作[]T，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已。

每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。

类型 []T 表示一个元素类型为 T 的切片。
*/

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	/**
	  切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

	  a[low : high]
	  它会选择一个半开区间，包括第一个元素，但排除最后一个元素。

	  以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：

	  a[1:4]
	*/
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	/**
	在进行切片时，你可以利用它的默认行为来忽略上下界。

	切片下界的默认值为 0，上界则是该切片的长度。

	对于数组

	var a [10]int
	来说，以下切片是等价的：

	a[0:10]
	a[:10]
	a[0:]
	a[:]

	*/

	s02 := []int{2, 3, 5, 7, 11, 13}

	s02 = s02[1:4]
	fmt.Println(s02)

	s02 = s02[:2]
	fmt.Println(s02)

	s02 = s02[1:]
	fmt.Println(s02)

	/**
	切片拥有 长度 和 容量。

	切片的长度就是它所包含的元素个数。

	切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

	切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
	*/

	s03 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s03 = s03[:0]
	printSlice(s03)

	// 拓展其长度
	s03 = s03[:4]
	printSlice(s03)

	// 舍弃前两个值
	s03 = s03[2:]
	printSlice(s03)

	/**
	切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。

	make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

	a := make([]int, 5)  // len(a)=5
	*/

	//a := make([]int, 5)
	//printSlice("a", a)
	//
	//b := make([]int, 0, 5)
	//printSlice("b", b)
	//
	//c := b[:2]
	//printSlice01("c", c)
	//
	//d := c[2:5]
	//printSlice01("d", d)
}

func printSlice01(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
