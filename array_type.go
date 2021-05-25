package main

import "fmt"

func main() {
	/*
		数据类型：
			基本类型：int，float，string，bool
			复合类型：array，slice，map，function，pointer，channel。。

		数组的数据类型：
			[size]type

		值类型：理解为存储的数值本身
			将数据传递给其他的变量，传递的是数据的副本（备份）
		引用类型：理解为存储的数据的内存地址
			slice，map
	*/

	//1.数据类型
	num := 10
	fmt.Printf("%T\n", num)

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [2]string{"hello", "world"}
	arr3 := [4]float64{1.222, 2.33, 4.33, 4.33}
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)
	fmt.Printf("%T\n", arr3)

	// 赋值
	num2 := num
	fmt.Println(num, num2)
	num2 = 20
	fmt.Println(num, num2)

	// 数组赋值
	arr4 := arr1
	fmt.Println(arr1)
	fmt.Println(arr4)

	arr4[0] = 100
	fmt.Println(arr1)
	fmt.Println(arr4)

	// 数值对比
	a := 3
	b := 4
	fmt.Println(a == b)
	fmt.Println(arr1 == arr4)
	// fmt.Println(arr1 == arr3) //  invalid operation: arr1 == arr3 (mismatched types [4]int and [4]float64)
}
