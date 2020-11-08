package main

import "image/color"

/**
通过嵌入结构体来扩展类型
 */

type ColoredPoint struct {
	Point
	Color color.RGBA
}

