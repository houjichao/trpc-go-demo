package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

/**
将Point结构体内嵌至ColoredPoint结构体中，从而实现“继承”（组合），并且ColoredPoint可以直接访问Point中的字段：
 */
func main()  {
	cp := ColoredPoint{}
	cp.X = 1
	fmt.Println(cp.Point.X)
	fmt.Println(cp.X)
}