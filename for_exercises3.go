package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		水仙花数：三位数：[100,999]
			没个位上的数字的立方和，刚好等于该数字本身，那么就叫做水仙花数，4个
			比如：153
				1*1*1+5*5*%+3*3*3=1+125+27=153
	*/
	for i := 100; i < 1000; i++ {
		x := i / 100
		y := i / 10 % 10
		z := i % 10
		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			fmt.Println(i)
		}
	}
	/*
		百位：1-9
		十位：0-9
		个位：0-9
	*/
	for x := 1; x < 10; x++ {
		for y := 0; y < 10; y++ {
			for z := 0; z < 10; z++ {
				n := x*100 + y*10 + z
				if x*x*x+y*y*y+z*z*z == n {
					fmt.Println(n)
				}
			}
		}
	}
}
