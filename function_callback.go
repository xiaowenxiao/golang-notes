package main

import "fmt"

func main() {
	/*
		高阶函数：
			根据go语言的数据类型的特点，可以将一个函数作为另一个函数的参数

		fun1(),fun2()
		将fun1函数作为了fun2这个函数的参数
			fun2函数：就叫做高阶函数
				接收了一个函数作为参数的函数，高阶函数
			fun1函数：回调函数
				作为另一个函数的参数的函数，叫做回调函数

	*/
	// 设计一个函数，用于求两个整数的加减乘除运算
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", oper)

	res1 := add(1, 2)
	fmt.Println(res1)

	res2 := oper(10, 20, add)
	fmt.Println(res2)

	res3 := oper(30, 10, sub)
	fmt.Println(res3)

	// 乘法(使用匿名函数作为回调函数)
	fun1 := func(a, b int) int {
		return a * b
	}
	res4 := oper(100, 3, fun1)
	fmt.Println(res4)

	// 除法(使用匿名函数作为回调函数)
	res5 := oper(100, 22, func(a, b int) int {
		if b == 0 {
			fmt.Println("除数不能为零")
			return 0
		}
		return a / b
	})
	fmt.Println(res5)
}

// 加法
func add(a, b int) int {
	return a + b
}

// 减法
func sub(a, b int) int {
	return a - b
}

func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b)
	return res
}
