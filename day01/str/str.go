package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	s_byte := []byte(s)
	fmt.Printf("%v\n%v\n%v", s, s_rune, s_byte)
	//返回值
	/*
		中国世界第一
		[20013 22269 19990 30028 31532 19968] s_rune的返回，可以是很大的数
		[228 184 173 229 155 189 228 184 150 231 149 140 231 172 172 228 184 128]  s_byte  表示的流的最大值不能超过255

	*/
	//查看rune和byte的数据类型
	fmt.Println()
	fmt.Println("~~~~~~~")
	var name string = "I'm 超人"
	a := []rune(name)
	b := []byte(name)
	fmt.Printf("%d:%T\n", a, a)
	fmt.Printf("%d:%T\n", b, b)
	/*
		[73 39 109 32 36229 20154]:[]int32
		[73 39 109 32 232 182 133 228 186 186]:[]uint8
	*/
	// 字符串拼接
	abc := "你好，我是："
	named := "超人"
	fmt.Println(abc+"1111", named)
	fmt.Printf("%s%s\n", abc, named)
	// 分片，以某个字段分片，如下
	var num string = "123-456-789"
	var num1 = strings.Split(num, "-")
	fmt.Println(num1)
	// 打印结果 [123 456 789]
	//strings.join()
	num2 := strings.Join(num1, "*")
	println(num2) // 返回值为 123*456*789
	// 字符串遍历
	var str6 = "hello 张三是个牛逼的程序员"
	//第一种遍历方式,for循环出来是byte类型
	fmt.Println("------------------")
	for i := 0; i < len(str6); i++ {
		fmt.Printf("%v=>", str6[i])
	}
	//第二种遍历方式, for循环出来是rune类型
	fmt.Println("------------------")
	//range循环，第一个返回值是index或key，第二个是value
	for _, value := range str6 {
		fmt.Printf("%T:%c\n", value, value)
	}
	//返回值
	/*
		104=>101=>108=>108=>111=>32=>229=>188=>160=>228=>184=>137=>230=>152=>175=>228=>184=>170=>231=>137=>155=>233=>128=>188=>231=>154=>132=>231=>168=>139=>229=>186=>143=>229=>145=>152=>------------------
		int32:hint32:eint32:lint32:lint32:oint32: int32:张int32:三int32:是int32:个int32:牛int32:逼int32:的int32:程int32:序int32:员
		进程 已完成，退出代码为 0

	*/
	//string转换，需要导入 strconv
	// int to string
	num3 := 200
	snum1 := strconv.Itoa(num3)
	fmt.Printf("%T:%v\n", snum1, snum1) //返回值 string:200
	// sting to int
	snum2, _ := strconv.Atoi(snum1)     //这个方法需要两个参数， 我们使用"_" 来声明，不使用，只是占用
	fmt.Printf("%T:%v\n", snum2, snum2) //返回值int:200

}
