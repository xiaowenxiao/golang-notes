package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "打开文件成功。。")
}
