package main //入口的库都为main
import "fmt"

/*
各种print 的区别
Println:
一次输入多个值的时候 Println 中间有空格
Println 会自动换行，Print 不会 Print:
一次输入多个值的时候 Print 没有 中间有空格
Print 不会自动换行 Printf
Printf 是格式化输出，在很多场景下比 Println 更方便 Sprintf
*/
//变量： 可以被调用和更改的值， 注意， 变量后面必须被调用，否则会报错
//声明全局变量
var name string = "erick"

//声明常量,常量是不可以更改的， 常量可以不被引用， 而变量一定要被后面引用，否则会报错
const abc = "abc"

// 批量声明多个常量
const (
	pi        = 3.14
	HOST_NAME = "weblogic"
)

//创建一个函数
func main() {
	fmt.Println(pi, HOST_NAME)
	//声明内部变量
	var gender = "male"
	salary := 20000 // 此种简写方法只能在fuc 内部使用
	//批量声明变量
	var d, e string
	d = "zhangsan"
	e = "18"
	//批量声明变量，第二种方法
	var (
		age  int
		name string
		c    bool
	)
	age = 18
	name = "erick"
	c = true
	//

	fmt.Println(d, e, c)
	fmt.Println(name, age, gender, salary)
	//格式化打印
	fmt.Printf("%s:%d\n%s:%v\n", name, age, gender, salary)
	//如果想使用字符串拼接，则可以使用
	a := fmt.Sprintf("name:%s\nage:%d\ngender:%s\nsalary:%v\n", name, age, gender, salary)
	fmt.Print("个人信息：\n", a)
}
