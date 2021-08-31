package main

import "fmt"

type Person struct {
	name string
}

func (p Person) show() {
	fmt.Println("Pseson--->", p.name)
}

// 类型别名
type People = Person

func (p Person) show2() {
	fmt.Println("People--->", p.name)
}

type Student struct {
	// 嵌入2个结构体
	Person
	People
}

func main() {
	var s Student
	// s.name ="张三" //ambiguous selector s.name
	s.Person.name = "张三"
	// s.show() //ambiguous selector s.show
	s.Person.show()

	fmt.Printf("%T,%T\n", s.Person, s.People)
	s.People.name = "小花"
	s.People.show2()

}
