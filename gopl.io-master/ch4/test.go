package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	// 数组
	var arrays [4]int
	fmt.Println(arrays[3])
	var arrays2 [4]string
	fmt.Println(arrays2[2])

	// 常量
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBD
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "&", GBD: "@", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])

	// map结构
	var makemap = make(map[string]string)
	makemap["test1"] = "1"
	makemap["test2"] = "2"
	for key, value := range makemap {
		fmt.Println("key:" + key + " value:" + value)
	}

	a, b := makemap["test3"]
	fmt.Println(a, "--", b) // 输出不存在的key, b的值为false，可作为判断有没有当前key值得依据

	// 匿名结构
	var point = Point{X: 3, Y: 5}
	var circle = Circle{Radius: 30, Point: point}
	circle.Point.Y = 66
	var wheel = Wheel{
		Circle: Circle{
			Point:  Point{X: 1, Y: 2},
			Radius: 3, // 换行了，最后就需要一个,否则编译不过
		}, Spoke: 20} // 同一行可以省略最后的,
	wheel.Y = 33
	fmt.Printf("circle: %#v \nwheel: %#v \n", circle, wheel)

	// JSON
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	var data, err = json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("json marshalling failed: %s\n", err)
	}
	//将byte[]转换为string
	fmt.Println("data to string is:", string(data[0:]))
	// json反编译
	var title []struct{ Title string }
	var error2 = json.Unmarshal(data, &title) // 这里传入的指针
	if error2 != nil {
		log.Fatalf("json marshalling failed: %s\n", error2)
	}
	fmt.Println("title is:", title)
	// 复杂一点的应用，请参考issue包对应的测试用例

	// 模板的使用，请参考issusereport包对应的测试用例
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spoke int
}
