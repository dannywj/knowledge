package main

/*
go 解析http请求返回的JSON 示例

json需要预先定义好对应的结构体，其中的属性可以不都定义，但是定义好的字段，数据类型必须一致
字段名不区分大小写，数据类型严格区分
go 结构体定义好的字段，json中可以不包含该字段
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 预先定义json对应的结构体
type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

func (s *Student) ShowStu() {
	fmt.Println("show Student :")
	fmt.Println("\tName\t:", s.Name)
	fmt.Println("\tAge\t:", s.Age)
	fmt.Println("\tGuake\t:", s.Guake)
	fmt.Println("\tPrice\t:", s.Price)
	fmt.Printf("\tClasses\t: ")
	for _, a := range s.Classes {
		fmt.Printf("%s ", a)
	}
	fmt.Println("")
}

func main() {
	println("======begin======")
	strData := httpGet()
	// 获取http get请求的json数据
	fmt.Println("http json result:" + string(strData))
	stb := &Student{}
	// json解析
	err := json.Unmarshal([]byte(strData), &stb)
	if err != nil {
		fmt.Println("Unmarshal faild")
	} else {
		// 解析成功，打印对象信息
		fmt.Println("Unmarshal success")
		stb.ShowStu()
	}

	println("======end======")
}

func httpGet() []byte {
	resp, err := http.Get("http://api.miyabaobei.com/echo.php")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	//fmt.Println(string(body))
	return (body)
}

/*
result show：

======begin======
http json result:{"name":"Dannywang","age":18,"Price":12.45,"Classes":["classA","classB"],"test":"test"}
Unmarshal success
show Student :
        Name    : Dannywang
        Age     : 18
        Guake   : false
        Price   : 12.45
        Classes : classA classB
======end======
*/
