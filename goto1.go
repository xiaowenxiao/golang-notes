package main

import "fmt"

func main() {
	var a = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}

		fmt.Printf("a的值为：%d\n", a)
		a++
	}

	fmt.Println("-------------------------")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakHere
			}
		}
	}
	return
breakHere:
	fmt.Println("done...")
}