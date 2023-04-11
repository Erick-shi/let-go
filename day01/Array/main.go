package main

import (
	"fmt"
)

func main() {
	//数字声明
	var arr = []int{1, 2, 3}
	fmt.Println(arr) //[1 2 3]
	//修改数值名称
	arr[0] = 10
	fmt.Println(arr) //[10 2 3]
	//数组遍历,第一种方法
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%T,%v\n", arr[i], arr[i])
		//第二种方法
		for _, v := range arr {
			fmt.Printf("%T,%v\n", v, v)
		}

	}
}
