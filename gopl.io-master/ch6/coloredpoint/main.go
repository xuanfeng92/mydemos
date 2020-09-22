// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 161.

// Coloredpoint demonstrates struct embedding.
package main

import (
	"fmt"
	"math"
)

//!+decl
import "image/color"

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}
type ColoredPoint2 struct {
	Point1 Point
	Point2 Point
	Color  color.RGBA
}

func (p *Point) updatePoint(newPoint Point) {
	p.Y = newPoint.Y
	p.X = newPoint.X
}
func (p Point) updatePoint2(newPoint Point) {
	p.Y = newPoint.Y
	p.X = newPoint.X
}

//!-decl

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	//!+main
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
	//!-main
	var p2 = ColoredPoint2{Point1: Point{X: 1, Y: 1}, Point2: Point{X: 5, Y: 4}, Color: red}
	var distance2 = p2.Point1.Distance(p2.Point2)
	fmt.Println("ps distance is:", distance2)
	// 测试方法更新对象
	var p1 = Point{X: 0, Y: 1}
	var newp1 = Point{X: 10, Y: 11}
	p1.updatePoint2(newp1)
	fmt.Printf("failed to update p1 value is:%v\n", p1)
	(&p1).updatePoint(newp1) // 通过绑定指针类型的方法，来更新数据
	fmt.Printf("success to update p1 value is:%v\n", p1)

}

/*
//!+error
	p.Distance(q) // compile error: cannot use q (ColoredPoint) as Point
//!-error
*/

func init() {
	//!+methodexpr
	p := Point{1, 2}
	//q := Point{4, 6}
	//
	//distance := Point.Distance   // method expression
	//fmt.Println(distance(p, q))  // "5"
	//fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy
	scale(&p, 2)
	//fmt.Println(p)            // "{2 4}"
	//fmt.Printf("%T\n", scale) // "func(*Point, float64)"
	//!-methodexpr
}

func init() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	//!+indirect
	type ColoredPoint struct {
		*Point
		Color color.RGBA
	}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) // "5"
	q.Point = p.Point                 // p and q now share the same Point
	p.ScaleBy(2)
	//fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
	//!-indirect
}
