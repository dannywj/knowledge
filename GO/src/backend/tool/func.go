package tool

import (
	"fmt"
	"reflect"
)

// 反射调用方法
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

// 计算整型绝对值
func CalcIntAbs(a int) (ret int) {
	ret = (a ^ a>>31) - a>>31
	return
}
