package main

import "fmt"

func main() {
	/*
		结构体：是有一系列具有相同类型或不同类型的数据构成的数据集合
			结构体成员是有一系列成员变量构成，这些成员变量也被成为"字段"
	*/
	// 1.方法一
	var p1 Person
	fmt.Println(p1)
	p1.name = "张三"
	p1.age = 15
	p1.sex = "男"
	p1.address = "深圳"
	fmt.Println(p1)
	fmt.Printf("姓名：%s,年龄：%d,性别：%s,地址：%s\n", p1.name, p1.age, p1.sex, p1.address)

	// 2.方法二
	p2 := Person{}
	p2.name = "Evan"
	p2.age = 20
	p2.sex = "男"
	p2.address = "曲靖"
	fmt.Printf("姓名：%s,年龄：%d,性别：%s,地址：%s\n", p2.name, p2.age, p2.sex, p2.address)

	// 3.方法三
	p3 := Person{name: "李四", age: 18, sex: "男", address: "上海"}
	fmt.Printf("姓名：%s,年龄：%d,性别：%s,地址：%s\n", p3.name, p3.age, p3.sex, p3.address)
	p4 := Person{
		name:    "王二",
		age:     26,
		sex:     "男",
		address: "西安",
	}
	fmt.Println(p4)

	// 4.

}

// 1.定义结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}
