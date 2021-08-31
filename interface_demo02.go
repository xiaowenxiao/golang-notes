package main

import "fmt"

func main() {
	/*
		空接口(interface{})
			不包含任何的方法，正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值

		fmt包下的Print系列函数：
			func Print(a ...interface{}) (n int,err error)
			func Printf(format string,a ...interface{}) (n int, err error)
			func Println(a ...interface{}) (n int, err error)
	*/
	var a1 A = Cat{"花猫"}
	var a2 A = Person{"王二", 30}
	var a3 A = "xixi"
	var a4 A = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	test1(a1)
	test1(a2)
	test1(3.14)
	test1("evan")

	test2(a3)
	test2(1000)

	// map,key字符串，value任意值
	map1 := make(map[string]interface{})
	map1["name"] = "诸葛黑蛋"
	map1["age"] = 30
	map1["friend"] = Person{"jelly", 30}
	fmt.Println(map1)

	// 切片，存储任意类型的数据
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, 100, "abc")
	fmt.Println(slice1)

	test3(slice1)
}
func test3(slice2 []interface{}) {
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("第%d个数据：%v\n", i+1, slice2[i])
	}
}

// 空接口
func test1(a A) {
	fmt.Println(a)
}

func test2(a interface{}) {
	fmt.Println("-->", a)
}

type A interface {
}

type Cat struct {
	color string
}

type Person struct {
	name string
	age  int
}
