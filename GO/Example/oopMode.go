// oopMode.go
package main

import "fmt"

// 学生结构体
type student struct {
	name string
	age  int
}

// 群组结构体（包含学生成员）
type group struct {
	student
	gname string
}

// 学生的方法
func (s student) sayHello() {
	fmt.Println("Hello student! my name is:" + s.name)
}

// 学生的方法（修改属性，传指针）
func (s *student) changeName(newname string) {
	s.name = newname
}

// 群组方法，覆盖了学生基类的方法
func (g group) sayHello() {
	fmt.Println("Hello group! my name is:" + g.name)
}

func main() {
	println("=======begin student list======")
	var s1 student
	s1.name = "wangjue"
	s1.age = 28

	fmt.Println("The student's name is: " + s1.name)

	// 派生类可以访问基类的属性
	var g1 group
	g1.gname = "danny group"
	g1.student.name = "g s name"
	fmt.Println("The group's name is: " + g1.gname + " sname:" + g1.student.name)

	println("--begin oo--")
	s1.sayHello()
	println("--begin change name--")
	s1.changeName("danny wang")
	println("--finish change name--")
	fmt.Println("The student's name is: " + s1.name)

	g1.sayHello()
	println("=======end student list======")
}
