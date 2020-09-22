package main

import (
	"./geometry"
	"fmt"
)

func main() {
	// 方法的使用
	var point1 = geometry.Point{X: 2, Y: 3}
	var point2 = geometry.Point{X: 5, Y: 8}
	var points = geometry.Path{point1, point2}
	var distance = points.Distance() // 定义任意类型的方法
	fmt.Println("distance is :", distance)

}
