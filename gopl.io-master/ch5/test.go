package main

import "fmt"

func main() {
	fmt.Println(addTest(1, 2))
	fmt.Println(addTest1(3, 7))
	fmt.Printf("%T\n", addTest)
	fmt.Printf("%T\n", addTest1)

	var pointTest = point{X: 1, Y: 2}
	addPoint(&pointTest) //如果想要改变传入实参的值，则需要定义指针类型的参数
	fmt.Println("y:", pointTest.Y)
}

type point struct {
	X int
	Y int
}

func addTest(a, b int) int {
	return a + b
}
func addTest1(a, b int) int {
	return a - b
}
func addPoint(p *point) {
	p.Y = p.X + p.Y
}
