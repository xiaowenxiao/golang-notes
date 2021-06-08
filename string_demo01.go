package main

import "fmt"

func main() {
	/*
		Go中的字符串是一个字节的切片
			可以通过将其内容封装在""中来创建字符串。go中的字符串是unicode兼容的，并且是UTF-8编码的。

		字符串是一些字节的集合。
			理解为一个字符的序列
			每个字符都有固定的位置(索引，下标，index：从0开始，到长度减1)

		语法："",``
			""
			"a","b","中"
			"abc","hello"
		字符：--->对应编码表中的编码值
			A-->65
			B-->66
			a-->97
			...
		字节：byte-->uint8
			utf8
	*/
	// 1.定义字符串
	s1 := "hello 中国" //一个汉字占用3个字节
	s2 := `hello world`
	fmt.Println(s1)
	fmt.Println(s2)

	// 2.字符串的长度:返回的是字节的个数
	fmt.Println(len(s1))
	fmt.Println(len(s2))

	// 3.获取某个字节
	fmt.Println(s2[0])
	a := 'h'
	b := 104
	fmt.Printf("%c,%c,%c\n", s1[0], a, b)

	// 4.字符串的遍历
	for i := 0; i < len(s2); i++ {
		// fmt.Println(s2[i])
		fmt.Printf("%c\t", s2[i])
	}
	fmt.Println()
	// for range
	for _, v := range s2 {
		fmt.Printf("%c\t", v)
	}
	fmt.Println()
	// 5.字符串是字节的集合
	slice1 := []byte{65, 66, 67, 68, 69}
	s3 := string(slice1)
	fmt.Println(s3)

	s4 := "abcdef"
	slice2 := []byte(s4)
	fmt.Println(slice2)

	// 6.字符串不能修改
	fmt.Println(s4)
	// s4[2] = `B`  //cannot assign to s4[2] (strings are immutable)
}
