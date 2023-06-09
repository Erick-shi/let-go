package main

import "fmt"

/*

Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，
将延迟处理的语句按defer定义的逆序进行执行，也就是说,
先被defer的语句最后被执行，最后被defer的语句，最先被执行。
*/
func main() {
	fmt.Println("1")
	defer fmt.Println("0")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
}

/*
打印结果
1
4
3
2
0
//会先执行自然打印的，然后在打印defer的内容，按照倒序来打印

*/
