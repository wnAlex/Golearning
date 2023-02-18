package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

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

type stringStruct struct {
	str unsafe.Pointer //str首地址
	len int            //str长度
}

func test10() {
	//for循环
	//1.基本写法
	for i := 1; i < 10; i++ {
		fmt.Println("hahaha", i)
	}

	//2.省略写法
	j := 0
	for j < 10 {
		fmt.Println("jjjj", j)
		j++
	}

	//3.死循环写法
	k := 0
	for {
		fmt.Println("kkkkk", k)
		k++
		if k > 10 {
			break
		}
	}
	fmt.Println()

	//遍历字符串
	s := "Alex王楠"
	fmt.Println(len(s))
	str := *(*stringStruct)(unsafe.Pointer(&s))
	fmt.Println(str.len)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d:%c\n", i, s[i])
	}
	fmt.Println()

	for i := 0; i < len([]rune(s)); i++ {
		fmt.Printf("%d:%c\n", i, []rune(s)[i])
	}
	fmt.Println()

	for i, value := range s {
		fmt.Printf("%d:%c\n", i, value)
	}
	fmt.Println()

}
func test11() {
	//break正常使用
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(j)
			if j == 2 {
				break
			}
		}
	}

	//break配合标签使用
	//label1:
	for i := 0; i < 4; i++ {
	label2:
		for j := 0; j < 4; j++ {
			fmt.Println(j)
			if j == 2 {
				//break label1
				break label2
			}
		}
	}

	//break配合随机数使用
	rand.Seed(time.Now().Unix())
	count := 0
	for {
		n := rand.Intn(100) + 1
		count++
		fmt.Println(n)
		if n == 99 {
			break
		}
	}
	fmt.Println(count)
}
func test12() {
	//continue正常使用
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				continue
			}
			fmt.Println(j)
		}
		fmt.Println()
	}
	//continue配合标签使用
label2:
	for i := 0; i < 4; i++ {
		//label1:
		for j := 0; j < 4; j++ {
			if j == 2 {
				//continue label1
				continue label2
			}
			fmt.Println(j)
		}
		fmt.Println()
	}
}

func test13() {
	for i := 0; i < 4; i++ {

		for j := 0; j < 4; j++ {
			if j == 2 {
				goto there
			}
			fmt.Println(j)
		}
		fmt.Println()
	}
there:
	fmt.Println("Im here!!")
}
func main() {
	//分支语句
	//test09()

	//循环语句
	//test10()

	//break 语句
	//test11()

	//continue语句
	//test12()

	//goto语句
	//test13()
}
