package main

import "fmt"

func main() {
	/*
		数据类型：
			一：按照数据类型来分：
				基本数据类型：
					int，float，string，bool
				复合数据类型：
					array，slice，map，struct，interface。。。
			二：按照数据的存储特点来分：
				值类型的数据：操作的是数值本身
					int，float64，bool，string，array
				引用类型的数据：操作的是数据的地址
					slice，map，chan
		参数传递：
			A：值传递
				传递的是数据的副本。修改数据，对于原始的数据没有影响
				值类型的数据，默认都是值传递，array，struct

			B：引用传递
				传递的是数据的地址，导致多个变量指向同一块内存
				引用类型的数据，默认都是引用传递，slice，map，chan

	*/
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前，数组的数据:", arr1)
	func1(arr1)
	fmt.Println("函数调用后，数组的数据：", arr1)

	arr3 := []int{6, 7, 8, 9}
	fmt.Println("函数调用前：切片的数据：", arr3)
	func2(arr3)
	fmt.Println("函数调用后，切片的数据：", arr3)

}

func func1(arr2 [4]int) {
	fmt.Println("函数中，数组的数据：", arr2)
	arr2[0] = 100
	fmt.Println("函数中，数组的数据修改后：", arr2)
}

func func2(arr4 []int) {
	fmt.Println("函数中，切片的数据：", arr4)
	arr4[0] = 999
	fmt.Println("函数中，切片的数据修改后：", arr4)
}
