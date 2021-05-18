package main

import "fmt"

func main() {
	/*
		题目1:
		*****
		*****
		*****
		*****
		*****

		题目2:
		1X1=1
		2X1=2 2X2=4
		3X1=3 3X2=6 3X3=9
		...

	*/
	for i := 0; i < 5; i++ {
		for i := 0; i < 5; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

}
