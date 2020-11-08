package main

type Abser interface {
	Abs() float64
}

func main() {
	//var a Abser
	//f := MyFloat(-math.Sqrt2)
	//v := Vertex{3, 4}
	//
	//a = f  // a MyFloat 实现了 Abser
	//a = &v // a *Vertex 实现了 Abser
	//
	//// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	//// 所以没有实现 Abser。
	//a = v
	//
	//fmt.Println(a.Abs())
}
