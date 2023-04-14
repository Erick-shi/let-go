package main

import "fmt"

func main() {
	fmt.Println("111")

}

func main() {
	//for 循环
	x := 0
	for {
		if x <= 10 {
			fmt.Println(x)
		} else {
			break
		}
		x++
	}
	fmt.Println("\n-----------")
	// if 判断
	y := 0
	if y > 0 {
		fmt.Println("1")
	} else if y < 0 {
		fmt.Println("2")
	} else {
		fmt.Println("3")
	}
	// break
	m := 0
	for {
		if m <= 5 {
			fmt.Println(m)

		} else {
			break
		}
		m++
	}
	fmt.Println("11111111111111111\n")
	// continue
	for x := 0; x < 10; x++ {
		if x == 5 {
			return
			println(x, "=====>")

		}
		fmt.Println(x)
	}
	// case
	score := "A"
	switch score {
	case "A":
		fmt.Println("nb")
	case "B":
		fmt.Println("jg")
	default:
		fmt.Println("1")
	}

}
