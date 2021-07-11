package main

import "fmt"

func main() {
	/*
		面向对象：OOP

		Go语言的结构体嵌套：
			1.模拟继承性：is - a
				type A struct{
					field
				}
				type B struct{
					A //匿名字段
				}
			2.模拟聚合关系：has - a
				type C struct{
					field
				}
				type D struct{
					c C //聚合关系
				}
	*/
	//  1.创建父类的对象
	p1 := Person{name: "张三", age: 30}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)

	// 创建子类的对象
	s1 := Student{Person: Person{name: "李四", age: 20}, school: "野鸡大学"}
	fmt.Println(s1)

	var s3 Student
	s3.Person.name = "王五"
	s3.Person.age = 40
	s3.school = "清华大学"
	fmt.Println(s3)

	s3.name = "evan"
	s3.age = 26
	fmt.Println(s3)

	fmt.Println(s1.name, s1.age, s1.school)
	fmt.Println(s3.name, s3.age, s3.school)
}

/*
s3.Person.name-->s3.name
Student结构体将Person结构体作为了一个匿名字段
那么Person中的字段，对于Student来讲，就是提升字段
Student对象直接访问Person中的字段
*/
// 1.定义父类
type Person struct {
	name string
	age  int
}

// 2.定义子类
type Student struct {
	Person        //模拟继承结构
	school string //子类的新增属性
}
