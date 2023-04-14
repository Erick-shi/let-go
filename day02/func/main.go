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
	x := a + b
	return x
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

}
