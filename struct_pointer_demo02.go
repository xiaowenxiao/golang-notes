package main

import "fmt"

func main() {
	/*
		数据类型：
			值类型：int，float，bool，string，array，struct
			引用类型：slice，map，function，pointer

		通过指针：
			new(),不是nil，空指针
				指向了新分配的类型的内存空间，里面存储的零值
	*/

	// 1.结构体是值类型
	p1 := Person{name: "王二狗", age: 20, sex: "男", address: "深圳市"}
	fmt.Println(p1)
	fmt.Printf("%p,%T\n", &p1, p1)

	p2 := p1
	fmt.Println(p2)
	fmt.Printf("%p,%T\n", &p2, p2)
	fmt.Println(p1)

	p2.name = "小杨"
	fmt.Println(p2)
	fmt.Printf("%p,%T\n", &p2, p2)

	//  2.定义结构体指针
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p,%T\n", pp1, pp1)

	// (pp1).name = "李四"
	pp1.name = "王五"
	fmt.Println(pp1)
	fmt.Println(p1)

	// 使用内置函数new()，go语言中专门用于创建某种类型的指针的函数
	pp2 := new(Person)
	pp2.sex = "女"
	pp2.address = "云南省"
	pp2.name = "yyl"
	pp2.age = 25
	fmt.Println(*pp2)

	// 创建任意类型的指针
	pp3 := new(int)
	fmt.Println(pp3)
	fmt.Println(*pp3)

}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}
