package main

import "fmt"

func main() {
	/*
		函数的参数：
			形式参数：也叫形参。函数定义的时候，用于接收外部传入数据的变量。
				函数中，某些变量的数值无法确定，需要由外部传入数据。
			实际参数：也叫实参。函数调用的时候，给形参赋值的实际的数。

			函数调用：
				1.函数名：声明的函数名和调用的函数名要统一
				2.实参必须严格匹配形参：顺序，个数，类型，一一对应的
	*/
	getSum(1)
	getSum(50)
	getSum(100)

	getSum1(12, 23)
	getSum2(22, 33)
	getSum3(1.22, 2.33, "hello")
}

func getSum(n int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1-%d的和：%d\n", n, sum)
}

func getSum1(a int, b int) {
	sum := a + b
	fmt.Printf("%d+%d=%d\n", a, b, sum)
}

func getSum2(a, b int) { //参数类型一致的可以简写在一起
	sum := a + b
	fmt.Printf("%d+%d=%d\n", a, b, sum)
}

func getSum3(a, b float64, c string) {
	fmt.Printf("a:%.2f,b:%.2f,c:%s\n", a, b, c)
}
