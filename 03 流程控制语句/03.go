package main

import "fmt"

func test09() {
	//分支语句
	//Golang 支持在if语句判断表达式中直接定义变量
	//a := 20
	//if i := a; i < 18 {
	//	fmt.Println(i)
	//} else {
	//	fmt.Println(i + 100)
	//}

	//bool类型不能与int类型混用了
	//b := 1
	//if b {      //错误
	//	fmt.Println("hahaha")
	//}

	//判读语句只能是关系表达式
	//b := true
	//if b=false {      //错误
	//	fmt.Println("hahaha")
	//}

	//switch-case 语句
	var e byte = 'e'
	fmt.Println(e)
	var d float64 = 2.2
	var a float64 = 2.2
	switch a {
	case 'a':
		fmt.Println("aaaa")
		fmt.Println("aaaa")
		fmt.Println("aaaa")
		fmt.Println("aaaa")
		fmt.Println("aaaa")
	case 30, 2.5, d:
		fmt.Println("bbbb")
		fallthrough
	case 'c':
		fmt.Println("cccc")
		fallthrough
	default:
		fmt.Println("no match")
	}

	switch dd := 20; {
	case d > 2.0:
		fmt.Println("dddd")
		fmt.Println(dd)
	case d < 2.0:
		fmt.Println("no dddd")

	}

	//type switch
	var x interface{} = 30
	switch y := x.(type) {
	case int32:
		fmt.Println("int32")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println(y)
	}
}
func test10() {
	//for循环
	for i := 1; i < 10; i++ {
		fmt.Println("hahaha", i)
	}

}
func main() {
	//分支语句
	//test09()

	//循环语句
	test10()
}
