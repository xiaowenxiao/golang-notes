package main

import "fmt"

func main() {
	// 第一种：定义变量，然后进行赋值
	var num1 int
	num1 = 30
	fmt.Printf("num1的数值是：%d\n", num1)
	// 写在一行
	var num2 int = 15
	fmt.Printf("num2的数值是：%d\n", num2)

	// 第二种：类型推断
	var name = "王二狗"
	fmt.Printf("类型是：%T，数值是：%s\n", name, name)

	// 第三种：简短定义也叫简短申明
	sum := 100
	fmt.Println(sum)

	// 多个变量同时定义
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	var m, n int = 100, 200
	fmt.Println(m, n)

	var n1, f1, s1 = 100, 3.14, "Go"
	fmt.Println(n1, f1, s1)

	var (
		studentName = "王小二"
		age         = 20
		sex         = "男"
	)
	fmt.Printf("学生姓名：%s,年龄：%d，性别：%s\n", studentName, age, sex)
}
