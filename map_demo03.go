package main

import "fmt"

func main() {
	/*
		map和slice的结合使用：
			1.创建map用于存储人的信息
				name，age，sex，address
			2.没个map存储一个人的信息
			3.将这些map存入到slice中
			4.打印遍历输出
	*/
	// 1.创建map存储第一个人的信息
	map1 := make(map[string]string)
	map1["name"] = "王二狗"
	map1["age"] = "30"
	map1["sex"] = "男"
	map1["address"] = "北京市"

	// 2.第二个人
	map2 := make(map[string]string)
	map2["name"] = "小花"
	map2["age"] = "20"
	map2["sex"] = "女"
	map2["address"] = "上海市"

	// 3.第三个人
	map3 := make(map[string]string)
	map3["name"] = "小张"
	map3["age"] = "26"
	map3["sex"] = "男"
	map3["address"] = "东莞市"

	// 将map存入到slice中
	s1 := make([]map[string]string, 0, 3)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)

	for i, val := range s1 {
		// val map1 map2 map3
		fmt.Printf("第 %d 个人的信息是：\n", i+1)
		fmt.Printf("\t姓名：%s\n", val["name"])
		fmt.Printf("\t年龄：%s\n", val["age"])
		fmt.Printf("\t性别：%s\n", val["sex"])
		fmt.Printf("\t地址：%s\n", val["address"])
	}
}
