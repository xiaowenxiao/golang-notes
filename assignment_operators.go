package main

import "fmt"

func main() {
	/*
		赋值运算符：
			=,+=,-=,*=,/=,%=,<<=,>>=,&=,|=,^=...
			=,把=右边的数值，赋值给=左边的变量

			+=，a += b ,相当于a = a + b
				a++,a += 1

			可以通过使用括号临时提升某个表达式的整体运算优先级

	*/
	var a int
	a = 3
	fmt.Println(a)

	a += 4
	fmt.Println(a)
	a -= 3
	fmt.Println(a)
	a *= 2
	fmt.Println(a)
	a /= 3
	fmt.Println(a)
	a %= 1
	fmt.Println(a)
}
