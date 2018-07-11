package main

/*
	反射调用示例
*/
import (
	//"errors"
	"fmt"
	"reflect"
)

func main() {
	// 声明一个map
	funcs := map[string]interface{}{
		"foo": foo,
		"bar": bar,
	}
	// 调用并获取返回值
	result := Call(funcs, "foo")
	fmt.Println(result[0].Interface())

	// 传递参数
	Call(funcs, "bar", 1, 2, 3)

	// 另外一个示例
	// fv := reflect.ValueOf(prints)
	// params := make([]reflect.Value,1)  //参数
	// params[0] := reflect.ValueOf(20)   //参数设置为20
	// rs := fv.Call(params)              //rs作为结果接受函数的返回值
	// fmt.Println("result:",rs[0].Interface().(string)) //当然也可以直接是rs[0].Interface()

}

func foo() int {
	// bla...bla...bla...
	fmt.Println("foo")
	return 123
}

func bar(a, b, c int) {
	// bla...bla...bla...
	fmt.Printf("bar %v,%v,%v\n", a, b, c)
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		fmt.Println("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
