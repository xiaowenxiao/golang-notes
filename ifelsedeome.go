package main

import "fmt"

func main() {
	/*
		if...else语句
			if 条件{
				//条件成立，执行此处的代码。。
				A段
			}else{
				//条件成立，执行此处的代码。。
				B段
			}
		注意点：
			1.if后的{,一定是和if条件写在同一行的
			2.else一定是if语句}之后，不能自己另起一行
	*/
	score := 60
	if score >= 60 {
		fmt.Println(score, "及格。。")
	} else {
		fmt.Println(score, "不及格。。")
	}

	if score >= 90 {
		fmt.Println(score, "优")
	} else if 60 <= score && score <= 90 {
		fmt.Println(score, "良")
	} else {
		fmt.Println(score, "差")
	}

}
