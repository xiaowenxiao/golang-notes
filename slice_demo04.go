package main

import "fmt"

func main() {
	/*
		按照类型来分：
			基本类型：int，float，string，bool
			复合类型：array，slice，map，struct，pointer，function，chan

		按照特点来分：
			值类型：int，float，string，bool，array
				传递的是数据副本
			引用类型：slice
				传递的地址，多个变量执行来统一快内存地址

		所以：切片是引用类型的数据，存储来底层数组的引用
	*/
	a1 := [4]int{1, 2, 3, 4}
	a2 := a1
	fmt.Println(a1, a2)
	a1[0] = 100
	fmt.Println(a1, a2)

	s1 := []int{1, 2, 3, 4}
	s2 := s1
	fmt.Println(s1, s2)
	s2[0] = 100
	fmt.Println(s1, s2)
	fmt.Printf("%p,%p\n", s1, s2)
	fmt.Printf("%p,%p\n", &s1, &s2)
}
