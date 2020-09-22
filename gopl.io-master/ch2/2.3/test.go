package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 1
	p := &x // p是整型指针，指向x
	fmt.Println(*p)
	fmt.Println(*p == x)           // *p表示p指针地址的值
	fmt.Println(reflect.TypeOf(p)) // p的类型的整型指针*int，指向x
	*p = 2
	fmt.Println(x)

	fmt.Println(f() == f())   // 返回的指针，每次不一样
	fmt.Println(*f() == *f()) // 但指针指向地址的值却是一样的

	v := 1
	incr(&v)              //这里递增后，相应的v变量的值也会递增，此时v=2
	fmt.Println(incr(&v)) // 此时v=3
	fmt.Println(v)        // 此时v=3

	v2 := new(int)
	v3 := *v2
	fmt.Println(reflect.TypeOf(v2))
	fmt.Println(v3)

	/*	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	circle()

}

func circle() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		y := x[i]
		if y != '!' {
			z := y + 'A' - 'a'
			fmt.Printf("%c", z)
		}
	}
	fmt.Println("\nlast result:" + x)
}
func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ //递增p所指向变量的值，但p自身不变
	return *p
}
