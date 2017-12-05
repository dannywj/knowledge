package main

/*
	go 解析 JSON 示例

	json需要预先定义好对应的结构体，其中的属性可以不都定义，但是定义好的字段，数据类型必须一致
	字段名不区分大小写，数据类型严格区分
	go 结构体定义好的字段，json中可以不包含该字段
	如果JSON中的字段在Go目标类型中不存在，json.Unmarshal() 函数在解码过程中会丢弃该字段。

	转换的字段必须是可导出的字段（type中定义的大写字母开头的，如果希望和str的字段不一致，可以使用tag方式改名）
*/
import (
	"encoding/json"
	"fmt"
)

// 预先定义json对应的结构体
type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32 `json:"ClassPrice"` //自定义导出字段
	hide    string  //小写字母字段不可导出json
	TestA   string
}

func (s Student) String() string {
	return fmt.Sprintf("student:[Name:%s,Age:%d,Price:%.2f,Classes:%s,Guake:%t]", s.Name, s.Age, s.Price, s.Classes, s.Guake)
}

func main() {
	println("======begin======")
	println("--[json -> obj]--")
	strData := "{\"name\":\"Dannywang\",\"age\":18,\"Price\":12.45,\"Classes\":[\"classA\",\"classB\"],\"test\":\"test\"}"
	fmt.Println("json str:" + string(strData))

	var stu Student
	// str转对象，需要传入地址
	err := json.Unmarshal([]byte(strData), &stu)
	if err != nil {
		fmt.Println("json error")
	}
	fmt.Println(stu)

	println("--[obj -> json]--")
	fmt.Println(stu)
	// 对象转str
	result, err := json.Marshal(stu)
	fmt.Println("json str:" + string(result)) //result为[]byte类型，需要转换为string

	// JSON与map的转换
	var map1, _ = Obj2map(stu)
	var map2, _ = JSON2map(strData)
	fmt.Println(map1)
	fmt.Println(map2)

	println("======end======")

}

func Obj2map(obj interface{}) (mapObj map[string]interface{}, err error) {
	// 结构体转json
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func JSON2map(jsonStr string) (mapObj map[string]interface{}, err error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

/*
result show：

======begin======
--[json -> obj]--
json str:{"name":"Dannywang","age":18,"Price":12.45,"Classes":["classA","classB"],"test":"test"}
student:[Name:Dannywang,Age:18,Price:0.00,Classes:[classA classB],Guake:false]
--[obj -> json]--
student:[Name:Dannywang,Age:18,Price:0.00,Classes:[classA classB],Guake:false]
json str:{"Name":"Dannywang","Age":18,"Guake":false,"Classes":["classA","classB"],"ClassPrice":0,"TestA":""}
map[Classes:[classA classB] ClassPrice:0 TestA: Name:Dannywang Age:18 Guake:false]
map[age:18 Price:12.45 Classes:[classA classB] test:test name:Dannywang]
======end======
*/
