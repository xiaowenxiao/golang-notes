package main

import "fmt"

func main() {
	/*
		深拷贝：拷贝的是数据本身
			值类型的数据，默认都是深拷贝：array，int，float，string，bool，struct

		浅拷贝：拷贝的是数据地址
			导致多个变量执行同一块内存
			引用类型的数据，默认都是浅拷贝：slice，map

			因为切换也是引用类型的数据，直接拷贝的是地址
		func copy(dst,src []Tyepe) int
			可以实现切换的拷贝
	*/
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0)
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)

	// copy()
	s3 := []int{7, 8, 9}
	fmt.Println(s2, s3)
	// copy(s2, s3)
	copy(s3[1:], s2[2:]) //将s2中的元素，拷贝到s3中
	fmt.Println(s2, s3)

}
