package main

import "fmt"

type Weekday int

const (
	firstday Weekday = iota
	sencondday
	thirthday
	forthday
)

func main() {
	var x uint8 = 1<<1 ^ 1<<5
	fmt.Printf("%b\n", x)
	fmt.Printf("%d\n", x)

	s := "hello世界"
	fmt.Printf("%d\n", len(s))
	fmt.Printf("%c\n", s[0])
	fmt.Printf("%s\n", s[2:])

	fmt.Println(thirthday)
}
