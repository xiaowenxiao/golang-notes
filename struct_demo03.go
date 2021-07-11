package main

import "fmt"

func main() {
	/*
		匿名结构体：没有名字的结构体
			在创建匿名结构体，同时创建对象
			变量名 := struct{
				定义字段Field
			}{
				字段进行赋值
			}

		匿名字段：一个结构体字段没有名字

		匿名函数：

	*/
	s1 := Student{name: "张三", age: 18}
	fmt.Println(s1.name, s1.age)
	func() {
		fmt.Println("Hello World")
	}()

	s2 := struct {
		name string
		age  int
	}{
		name: "李四",
		age:  20,
	}
	fmt.Println(s2.name, s2.age)

	w2 := Worker{"王二", 32}
	fmt.Println(w2)
	fmt.Println(w2.string)
	fmt.Println(w2.int)
}

type Worker struct {
	string //匿名字段
	int    //匿名字段，默认使用数据类型作为名字，只能使用一次
}

type Student struct {
	name string
	age  int
}
