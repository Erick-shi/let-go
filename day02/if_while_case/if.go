package main

import "fmt"

func main() {
	var x []int
	fmt.Println(x)
	var y = make([]int, 0)
	y = append(y, []int{1, 2, 3, 4}...)
	fmt.Printf("%p,%T,%v\n", y, y, y)
	ny := &y
	fmt.Printf("%p,%T,%v", ny, ny, ny)

}
