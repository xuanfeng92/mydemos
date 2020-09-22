package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	var w io.Writer
	fmt.Println(w == nil)
	w = os.Stdout
	w.Write([]byte("呵呵呵1"))
	w = new(bytes.Buffer)
	w.Write([]byte("hehehe"))
	//w = nil
	fmt.Println(w == nil)
	fmt.Printf("%s\n", w)

	// 验证nil
	const debug = true
	var buf *bytes.Buffer // 解决办法： 把*bytes.Buffer 换成io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	fmt.Println("is nil?", buf == nil)
	f(buf)
	fmt.Printf("%s\n", buf)

	// 类型断言
	var t io.Writer
	t = os.Stdout
	//f := t.(*os.File)
	fmt.Printf("type: %T", t)
	//c := t.(*bytes.Buffer) // 运行会报错： panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
	fmt.Println(f)

	var rw io.Writer
	rw = os.Stdout
	writer, ok := rw.(*bytes.Buffer)
	if ok {
		fmt.Println(reflect.TypeOf(writer))
	} else {
		fmt.Println(writer)
	}

	//writer := rw.(io.ReadWriter)

	fmt.Println(rw)

	// 测试类型分支
	fmt.Println(sqlQuote(nil))
	fmt.Println(sqlQuote(-2))
	fmt.Println(sqlQuote(23))
	fmt.Println(sqlQuote(true))
	fmt.Println(sqlQuote(false))
	fmt.Println(sqlQuote("where i =0"))
	fmt.Println(sqlQuote(os.Stdout))

}
func f(out io.Writer) {
	// 这里比较的时候，out类型必须是io.Writer或者实现了该接口的具体类型或者指明了是nil类型，否则会匹配不上
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

func sqlQuote(x interface{}) string {
	switch x.(type) { // 这里的type是关键字，表示类型分支
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x.(bool) {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return fmt.Sprintf("%s", x)
	default:
		panic(fmt.Sprintf("unexpected type %T:%v", x, x))
	}
}
