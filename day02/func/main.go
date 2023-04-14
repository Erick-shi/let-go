package main

import (
	"fmt"
)

//定义一个空函数：func name(){}
func Func1() {
	fmt.Printf("this is a null func\n")
}

//定义一个接收string 参数的函数： func Func2(name string){}
func Func2(name string) {
	fmt.Println(name)
}

//定义一个接收多值的函数
func Func3(a int, b int) int {
	x := a + b //注意这 必须使用定义变量的方法声明
	return x
}

//定义一个接收多值的函数，且有返回值
func Func4(a int, b int) (resp int) {
	resp = a + b //此时不用复变量，可以直接赋值的方式
	return resp
}

//不仅可以定义固定参数，还可以定义可变参数， 可变参数为不固定参数，，，下面定义一个把所有输入的参数求和的函数
func Func5(n ...int) int {
	sum := 0
	for _, v := range n {
		sum = sum + v
	}
	return sum
}

// 在一个函数中，可以同时接收可变参数和不可变参数，可变参数一定要放到后面！
func Func6(a int, b ...int) int {
	sum1 := a
	for _, v := range b {
		sum1 = sum1 + v
	}
	return sum1
}

//多个返回值
func Func7(x int, b int) (xx int, yy int) {
	xx = x + b
	yy = x - b
	return
}

func main() {
	Func1()
	// 调用Func2
	name := "erick"
	Func2(name)
	//或者直接调用
	Func2("我的中国薪")
	// 调用func3
	x := Func3(2, 10)
	fmt.Println(x)
	//调用Func4
	y := Func4(20, 30)
	fmt.Println(y)
	//调用Func5，求1+2+3+4+5的和
	xx := Func5(1, 2, 3, 4, 5)
	fmt.Println(xx)
	//调用func6
	yy := Func6(1)        // a=[1]
	yy1 := Func6(1, 2, 3) // a=[1] b=[2,3]
	fmt.Println(yy)
	fmt.Println(yy1)
	// 调用Func9
	m, n := Func7(100, 200)
	fmt.Println(m, n) //300 -100

}
