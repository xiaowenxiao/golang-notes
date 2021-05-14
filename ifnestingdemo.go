package main

import "fmt"

func main() {
	/*
		if语句的嵌套：
		if 条件1{
			A段
		}else{
			if 条件2{
				B段
			}else{
				C段
			}
		}

		简写：
		if 条件1{
			A段
		}else if 条件2{
			B段
		}else if 条件3{
			C段
		}...else{

		}
	*/
	sex := ""
	if sex == "男" {
		fmt.Println("可以去男厕所。。。")
	} else {
		if sex == "女" {
			fmt.Println("去女厕所吧。。。")
		} else {
			fmt.Println("我也不知道了。。。")
		}
	}

	if sex == "男" {
		fmt.Println("可以去男厕所。。。")
	} else if sex == "女" {
		fmt.Println("去女厕所吧。。。")
	} else {
		fmt.Println("我也不知道了。。。")
	}

}
