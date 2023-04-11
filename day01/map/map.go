package main

import "fmt"

/*
map 可以理解为py中的字典，即k:v格式的数据类型
*/
func main() {
	//define a map; map[key类型]v类型{}
	userInfo := map[string]string{
		"username":     "erick",
		"home address": "bj",
		"age":          "18",
	}
	fmt.Println(userInfo)
	//或者使用make定义
	map1 := make(map[string]int)
	//添加数据
	map1["age"] = 18
	map1["salary"] = 10000
	fmt.Println(map1)
	// 判断某个值是否存在
	v, ok := userInfo["username"]
	if ok {
		println(v)
	}
	// 删除map 或者key,用法delete(map 对象，key)
	delete(userInfo, "username")
	fmt.Println(userInfo)
	// map 遍历
	for k, v := range map1 {
		println(k, v)
	}
	// map 只遍历key
	for k := range map1 {
		println(k)
	}
	//指针
	a := 1
	fmt.Println(a)       //1   返回值
	fmt.Println(&a)      //0x1400001c0c0  指针地址
	fmt.Println(*&a)     // 1  指针取值
	fmt.Printf("%T", &a) //*int0x1400001c0c0  指针地址的类型， var p *int =  &a

	b := &a
	fmt.Println(b)   // a的指针地址
	fmt.Println(&b)  // b的指针地址
	fmt.Println(*&b) // b取值

	/*
		1.

	*/
}
