package main

import "fmt"

func main() {
	/*
		结构体嵌套：一个结构体中的字段，是另一个结构体类型
	*/
	b1 := Book{}
	b1.bookNmae = "水浒传"
	b1.price = 80.259
	fmt.Println(b1)

	s1 := Student{}
	s1.name = "王二狗"
	s1.age = 19
	s1.book = b1 //值传递
	fmt.Println(s1)
	fmt.Printf("学生姓名：%s,学生年龄：%d,看的书是：《%s》,书的价格是：%.2f\n", s1.name, s1.age, s1.book.bookNmae, s1.book.price)

	s1.book.bookNmae = "西游记"
	fmt.Println(s1)
	fmt.Println(b1)

	s2 := Student{name: "李小花", age: 19, book: Book{bookNmae: "《小白学go编程》", price: 89.3232}}
	fmt.Println(s2.name, s2.age)
	fmt.Println("\t", s2.book.bookNmae, s2.book.price)

	s3 := Student{
		name: "jelly",
		age:  25,
		book: Book{
			bookNmae: "《十万个为甚么》",
			price:    90.00,
		},
	}
	fmt.Println(s3.name, s3.age)
	fmt.Println("\t", s3.book.bookNmae, s3.book.price)

	b4 := Book{bookNmae: "呼啸山庄", price: 76.5}
	s4 := Student2{name: "Evan", age: 26, book: &b4}
	fmt.Println(b4)
	fmt.Println(s4)
	fmt.Println("\t", s4.book)

	s4.book.bookNmae = "北京往事"
	fmt.Println(b4)
	fmt.Println(s4)
	fmt.Println("\t", s4.book)

}

// 1.定义一个书的结构体
type Book struct {
	bookNmae string
	price    float64
}

// 2.定义学生的结构体
type Student struct {
	name string
	age  int
	book Book
}

type Student2 struct {
	name string
	age  int
	book *Book //book的地址
}
