package main

import (
	"fmt"
	"reflect"
)

type MyType struct {
	Type string
	Num  int
	Body myBody
}
type myBody struct {
	Name  string
	value string
}

func (b *myBody) SetValue(value string) {
	if len(value) > 0 {
		b.value = value
	}
}
func main() {
	v := reflect.ValueOf(1)
	fmt.Println("通过接口获取值：", v)
	i := v.Interface()
	fmt.Println("通过类型断言获取值：", i.(int))
	fmt.Println("通过反射获取类型Type:", reflect.TypeOf(struct {
		Type string
		Num  int
	}{Num: 25, Type: "type1"}))
	var a = MyType{
		Num:  20,
		Type: "type1",
		Body: myBody{
			Name:  "body111",
			value: "value111",
		},
	}
	b := reflect.ValueOf(&a).Elem() // 寻址 代表a变量，注意：里面&a 指向了a变量。
	fmt.Println("是否可以寻址：", b.CanAddr())
	bx := b.Addr().Interface().(*MyType) // 相当于 *a，这是由于Addr()会包含一个指针，因为b那里使用了&a
	fmt.Println("换算成指针：", bx)
	bx.Body.Name = "change222"
	fmt.Println("通过寻址改变内容：", a)

	body0 := b.FieldByName("Body")
	if body0.CanAddr() {
		body := body0.Addr().Interface().(*myBody) // 通过寻址改变子结构的类型
		body.value = "body中的value改变了"
		body.Name = "body中的Name 改变了"
		fmt.Println("---看看有没有改变----：", a)

		mNum := body0.NumMethod()
		fmt.Println("反射获取的方法数量：", mNum)
		if mNum > 0 {

			methodName := body0.Method(0)
			fmt.Println("第一个方法名为：", methodName)
		}

	}

	var privateFiled = b.FieldByName("Body").FieldByName("value")
	fmt.Println("获取未导出的字段:", privateFiled) // 可以遍历到未导出的类型
	fmt.Println("未导出的类型可以被addr到：", privateFiled.CanAddr())
	fmt.Println("未导出的类型不能被set：", privateFiled.CanSet())
	if privateFiled.CanSet() && privateFiled.CanAddr() {
		privateFiled.SetString("gogogo") // ！！！这里修改私有的字段，会报错：reflect.Value.SetString using value obtained using unexported field
	}

	body := b.FieldByName("Body").Interface().(myBody)
	body.value = "change112233"
	fmt.Println(b.FieldByName("Body").FieldByName("value"))
	var e = &a.Body
	e.value = "changge6666"
	fmt.Println(a)

	fmt.Println(b.NumField()) // 3
	var c = 2
	d := reflect.ValueOf(&c).Elem()
	fmt.Println(d.CanAddr())

	//c := b.Field(2).Field(1).Interface().(string)
	//fmt.Printf("hehe:%s", c)
	//fmt.Printf("%v", b.CanInterface())
	//fmt.Println(b)
	//fmt.Println(b.Type().Field(0).Name)

	fmt.Println()
}
