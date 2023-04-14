package main

import "fmt"

func main() {
	var list1 []string //声明一个空切片
	fmt.Println(list1) //返回值[]
	//第一种声明切片的方法, make 声明 make([]init(类型)，10 长度)注意这个10这个长度是可变的
	var x = make([]int, 0)
	x = append(x, 1, 2, 3, 4111)
	fmt.Println(x)
	x = append(x, []int{5, 6, 8}...)
	fmt.Println(x)
	//声明一个切片,并且赋值
	var list2 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(list2) // [1 2 3 4 5 6 7 8]
	// 截取前3个key
	fmt.Println(list2[:3]) //[1 2 3]  计数是从0开始， 前包，结尾不包。即下标为0-2
	//汲取从第2到最后
	fmt.Println(list2[2:]) //[3 4 5 6 7 8]
	fmt.Printf("%T:%v", nil, nil)
	//切片循环，第一种方法
	for i := 0; i < len(list2); i++ {
		println(list2[i])
	}
	//循环第二种方法
	for _, value := range list2 {
		println(value)
	}
	// 值替换
	list2[0] = 100
	println(list2[0])
	// 内存地址copy，相当于linux 中的软链接
	list3 := list2
	list3[1] = 101            //修改了list3， list2也会被修改，因为是同一块地址空间
	fmt.Println(list3, list2) //[100 101 3 4 5 6 7 8] [100 101 3 4 5 6 7 8]
	fmt.Printf("%T:%d\n%T:%d\n", list3, list3, list2, list2)
	fmt.Printf("%p:%d\n%p:%d\n", list3, list3, list2, list2) //两list 是同一块地址空间
	//切片 append 追加
	var list4 []int           //声明一个空的切片
	fmt.Println(list4 == nil) //true
	fmt.Println(list4)        //[]
	list4 = append(list4, 1)
	fmt.Println(list4) //[1]
	// for 循环append
	var list5 []int
	for i := 0; i < 10; i++ {
		list4 = append(list5, i)
		fmt.Printf("%T,%d", list5)
	}
	var list6 []string
	list6 = append(list6, "erick", "bj")
	fmt.Println(list6)
	// 删除切片中的元素
	list7 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%p\n", list7)
	list7 = append(list7[:2], list7[3:]...)
	fmt.Printf("%p\n", list7)
	fmt.Println(list7)
	//切片合并
	list9 := []int{1, 2, 3, 4}
	list10 := []int{5, 6, 7, 8, 9}
	//list12 := []int{10, 11, 12}
	list11 := append(list10, list9...)
	fmt.Println(list11)
}
