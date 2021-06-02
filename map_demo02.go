package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		map的遍历：
			使用：for range
				数组，切片：index，value
				map：key，value
	*/
	map1 := make(map[int]string)
	map1[1] = "峰哥"
	map1[2] = "亡命"
	map1[3] = "天涯"
	map1[4] = "卡秃噜皮"
	map1[5] = "君"
	map1[6] = "盼盼"

	// 1.遍历map
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	fmt.Println("_______________________")
	for i := 1; i <= len(map1); i++ {
		fmt.Println(i, "---->", map1[i])
	}

	/*
		1.获取所有的key，-->切片/数组
		2.进行排序
		3.遍历key,--->map[key]
	*/
	keys := make([]int, 0, len(map1))
	fmt.Println(keys)

	for k, _ := range map1 {
		keys = append(keys, k)
	}
	// 冒泡排序，或者使用sort包下的排序方法
	sort.Ints(keys)
	fmt.Println(keys)

	for _, v := range keys {
		fmt.Println(v, "-->", map1[v])
	}

	s1 := []string{"Apple", "Windows", "IOS", "abc", "王二狗", "acc"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)

}
