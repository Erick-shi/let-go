package main

import "fmt"

//err 是一个意料之中的判断，属于代码逻辑中的东西
/*
err 是一个意料之中的判断，属于代码逻辑中的东西
func main() {
	resp, err := http.Get("http://www.111111baidu.com")
	if err != nil {
		fmt.Println("this web is cat not visit")
		return
	} else {
		fmt.Println(resp.StatusCode)
	}
}

*/

//panic 异常是意外发生的，需要手动捕获。
func main() {
	fmt.Println("the start~~~~~~~")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获panic异常", r)
		}
	}()
	var x map[string]string
	fmt.Println(x)
	x["name"] = "zhangsan"        // 若不捕获则会抛出：panic: assignment to entry in nil map
	fmt.Println("the end~~~~~~~") //后面的将不会打印
}
