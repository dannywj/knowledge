package main

// 定义人的结构体【基础】
type person struct {
	name string
	age  int
}

// 定义学生的结构体
type student struct {
	no     string
	person // 匿名对象
}

// 定义教师的结构体
type teacher struct {
	level string
	person
}

// 人类的方法（基类方法）
func (p person) sayHello() {
	println("person:" + p.name + " say hello")
}

// 学生的方法 写作业
func (s student) doHomework() {
	println("student:" + s.name + " (no:" + s.no + ") can do homework")
}

// 老师的方法 上课
func (t teacher) takeLession() {
	println("teacher:" + t.name + " (level:" + t.level + ") can teaching")
}

// 学生的方法 吃饭
func (s student) eat() {
	println("student eat:" + s.name + " (no:" + s.no + ") ******")
}

// 老师的方法 吃饭
func (t teacher) eat() {
	println("teacher eat:" + t.name + " (level:" + t.level + ") ######")
}

// 定义接口（一组方法的集合，包含吃饭和打招呼）
type Men interface {
	eat()
	sayHello()
}

// 实现了String方法的对象，可以直接输出（重写String）
// fmt.Println(stu)
func (s student) String() string {
	return "student is:" + s.name
}

func main() {
	println("======begin======")
	var stu student
	stu.age = 18
	stu.name = "xiaoming"
	stu.no = "001"

	var tea teacher
	tea.age = 35
	tea.name = "Jack"
	tea.level = "normal"

	stu.doHomework()
	tea.takeLession()

	stu.eat()
	tea.eat()
	tea.sayHello()

	var m Men
	m = stu
	m.sayHello()
	println("======end======")
}
