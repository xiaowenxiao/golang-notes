package main

import "fmt"

func main() {
	/*
		switch中的break和fallthrough语句
		break：可以使用在switch，也可以使用在for循环中
			强制结束case语句，从而结束switch分支

		fallthrough：用于穿透switch
			当switch中的某个case匹配成功后，就执行该case语句
			如果遇到fullthough，那么后面紧邻的case，无需匹配，执行穿透

			fallthrough应该位于某个case的最后一行
	*/
	n := 2
	switch n {
	case 1:
		fmt.Println("为111")
		fmt.Println("为111")
		fmt.Println("为111")
	case 2:
		fmt.Println("为222")
		fmt.Println("为222")
		break
		fmt.Println("为222")
	case 3:
		fmt.Println("为333")
		fmt.Println("为333")
		fmt.Println("为333")
	}

	m := 3
	switch m {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
		fallthrough
	case 4:
		fmt.Println("第四季度")
	}
}
