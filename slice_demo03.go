package main

import "fmt"

func main() {
	/*
		slice := arr[start:end]
			切片的数据：[start:end]
			arr[:end],从头到尾
			arr[start:],从start到末尾

		从已有的数组上，直接创建切片，该切片的底层数组就是当前的数组。
			长度是从start到end切割的数据量。
			但是容量从start到数组的末尾。
	*/
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("-----------------1.已有数组直接创建切片-----------------")
	s1 := a[:5]
	s2 := a[3:8]
	s3 := a[5:]
	s4 := a[:]
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)
	fmt.Println("-----------------2.长度和容量-----------------")
	fmt.Printf("s1  len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("s1  len:%d,cap:%d\n", len(s2), cap(s2))
	fmt.Printf("s1  len:%d,cap:%d\n", len(s3), cap(s3))
	fmt.Printf("s1  len:%d,cap:%d\n", len(s4), cap(s4))

	fmt.Println("-----------------3.更改数组的内容-----------------")
	s1 = append(s1, 1, 1, 1, 1)
	fmt.Println(s1)
	fmt.Println(a)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Println("-----------------4.添加元素切片扩容-----------------")
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 2, 2, 2, 2)
	fmt.Println(s1)
	fmt.Println(a)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", &a)
}
