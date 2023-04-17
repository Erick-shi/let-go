package main

import "fmt"

//回顾匿名函数，匿名函数即没有名字的函数， 不可以随意被调用。下面定义一个匿名函数

/*
func main() {
	func() {
		fmt.Println("我是一个匿名函数")
	}()
}
*/
//闭包

/*
闭包为 函数加+外层变量引用.
func
*/
//函数里面套一个函数，当里面的函数需要变量时，首先会遍历自身函数的变量，若找不到则向外层遍历
func a() func() {
	name := "我是一个闭包"
	return func() {
		fmt.Println("hello", name)
	}
}

//我们把他变得更复杂些
func b(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}

}
func main() {
	n := a()
	n()                           // hello 我是一个闭包.相当于执行了a函数内部的匿名函数
	fmt.Printf("%p,%p\n", n, a()) //0x100ef91f0,0x100ef91f0
	//
	n1 := b("张三")
	n1()

	x3 := c()
	x3()
	x4 := d(18)
	x4()
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7

}

func c() func() {
	return func() {
		fmt.Println("我是你爸爸")
	}
}
func d(age int) func() {
	return func() {
		fmt.Println("我今年", age, "岁了")
	}
}

//
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub

}

//

//func e(base int) (func(int) int, func(int) int) {
//	add := func(i int) {
//		base = base + i
//		return base
//	}
//	sub := func(i int) {
//		base = base + i
//		return base
//	}
//	return add, sub
//}
