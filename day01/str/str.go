package main

import "fmt"

//定义带格式的字符串
func main() {
	s1 := `
	name:"erick"
	age: "18"
	salary: "2000"`
	fmt.Println(s1)
	/* 在go 中有两种 语言字符
	1. unit8 类型或者叫 byte 类型，代表ASCII码
	2. rune 类型： 也称之为 utf-8类型
	其实rune类型即utf-8类型的底层也是ASSII码，
	*/
	s := "中国世界第一"
	s_rune := []rune(s)
	fmt.Printf("%v\n%v", s, s_rune)
}
