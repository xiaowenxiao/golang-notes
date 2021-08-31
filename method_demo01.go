package main

import (
	"fmt"
)

func main() {
	/*
		方法：method
			一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者一个指针。
			所有给定类型的方法属于该类型的方法集

		语法：
			func (接受者) 方法名(参数列表)(f返回值列表){

			}

		总结：
			method，同函数类似，区别需要有接受者。(也就是调用者)

		对比函数：
			A：意义
				方法：某个类别的行为功能，需要指定的接收者调用
				函数：一段独立功能的代码，可以直接调用

			B：语法
				方法：方法名可以相同，只要接受者不同
				函数：命名不能冲突


	*/
	w1 := Worker{name: "张三", age: 18, sex: "男"}
	fmt.Printf("%T\n", w1)
	w1.work()

	w2 := &Worker{name: "edison", age: 40, sex: "男"}
	fmt.Printf("%T\n", w2)
	w2.work()

	w2.rest()
	w1.rest()

	w2.printInfo()
	c1 := Cat{color: "白色", age: 1}
	c1.printInfo()
}

// 2.定义一个工人结构体
type Worker struct {
	name string
	age  int
	sex  string
}

type Cat struct {
	color string
	age   int
}

// 2.定义行为方法
func (w Worker) work() {
	fmt.Println(w.name, "在工作。。。")
}

func (p *Worker) rest() {
	fmt.Println(p.name, "在休息。。。")
}

func (p *Worker) printInfo() {
	fmt.Printf("工人姓名：%s,工人年龄：%d,工人性别：%s\n", p.name, p.age, p.sex)
}

func (p *Cat) printInfo() {
	fmt.Printf("猫咪的颜色：%s,年龄：%d\n", p.color, p.age)
}
