package main

import (
	"fmt"
	"math"
)

func main() {
	var t1 Triangle = Triangle{3, 5, 6}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a, t1.b, t1.c)

	var c1 Circle = Circle{4}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())
	fmt.Println(c1.radius)

	var s1 Shape
	s1 = t1
	fmt.Println(s1.area())
	fmt.Println(s1.peri())

	var s2 Shape
	s2 = c1
	fmt.Println(s2.area())
	fmt.Println(s2.peri())

	testShape(t1)
	testShape(c1)
	testShape(s1)
	testShape(s2)

	getType(t1)
	getType(c1)
	getType(s1)

	var t2 *Triangle = &Triangle{3, 4, 2}
	fmt.Printf("t2:%T,%p,%p\n", t2, t2, &t2)
	getType(t2)

	getType2(t2)
	getType2(t1)
}

func getType2(s Shape) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形。。", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("圆形。。", ins.radius)
	case *Triangle:
		fmt.Println("三角形结构体指针：", ins.a, ins.b, ins.c)
	}
}

func getType(s Shape) {
	// 断言
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三边是：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径是：", ins.radius)
	} else if ins, ok := s.(*Triangle); ok {
		fmt.Printf("ins:%T,%p,%p\n", ins, ins, &ins)
		fmt.Printf("s:%T,%p,%p\n", s, s, &s)
	} else {
		fmt.Println("我也不知道了。。。")
	}
}

func testShape(s Shape) {
	fmt.Printf("周长：%.2f,面积：%.2f\n", s.peri(), s.area())
}

// 1.定义一个接口
type Shape interface {
	peri() float64 //形状的周长
	area() float64 //形状面积
}

// 2.定义实现类：三角形
type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

// 3.定义实现类：圆形
type Circle struct {
	radius float64
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}
func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
