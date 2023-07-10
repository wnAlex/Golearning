package main

import (
	"errors"
	"fmt"
	"github.com/wnAlex/Golearning/04functionAndPackage/myfunc"
	"github.com/wnAlex/Golearning/04functionAndPackage/myfunc2"
	"strings"
)

func test01() {
	n1 := 1.2
	n2 := 2.2
	n3 := "%"
	res, err := myfunc.Cal(n1, n2, n3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
func cal(n1 int, n2 int) *int {
	n3 := n1 + n2
	return &n3
}
func test02() {
	//由于逃逸分析的存在，可能返回局部变量的地址不会出错，但是不建议这么做
	res := cal(10, 20)
	fmt.Println(*res)
	fmt.Println(res)
	i := 100
	fmt.Println(i)
	fmt.Println(*res)
	fmt.Println(res)
	*res = 100
	fmt.Println(*res)
	fmt.Println(res)
}
func printArr(arr [4]int) {
	for index, _ := range arr {
		fmt.Printf("%d", arr[index])
		arr[index]++
		fmt.Printf("  %d", arr[index])
		fmt.Println()
	}
}
func test03() {
	arr := [4]int{1, 2, 3, 4}
	printArr(arr)
	for index, _ := range arr {
		fmt.Printf("%d ", arr[index])
	}
	fmt.Println()
}
func change(n1 float64, n2 float64, n3 *float64) {
	*n3 = n1 + n2
}
func test04() {
	n1 := 20.1
	n2 := 10.1
	var n3 float64 = 0
	fmt.Println(n3)
	change(n1, n2, &n3)
	fmt.Println(n3)
}
func change2(s1 []int) {
	for index, _ := range s1 {
		fmt.Printf("%d ", s1[index])
		s1[index]++
		fmt.Printf("after change %d ", s1[index])
		fmt.Println()
	}
}
func test05() {
	s1 := []int{1, 2, 3, 4}
	for index, _ := range s1 {
		fmt.Printf("%d ", s1[index])
	}
	fmt.Println()
	change2(s1)
	for index, _ := range s1 {
		fmt.Printf("%d ", s1[index])
	}
	fmt.Println()
}
func getAdd(n1 int, n2 int) int {
	return n1 + n2
}
func funcpara(para func(int, int) int, n1 int, n2 int) int {
	return para(n1, n2)
}
func funcreturn() func(int, int) int {
	return getAdd
}
func test06() {
	a := getAdd
	fmt.Printf("the type of a is %T\n", a)
	res := a(20, 30)
	fmt.Println(res)

	res2 := funcpara(getAdd, 50, 60)
	fmt.Println(res2)

	res3 := funcreturn()
	fmt.Println(res3(100, 200))

}

type myfunctype func(int, int) int

func funcpara2(para myfunctype, n1 int, n2 int) int {
	return para(n1, n2)
}
func test07() {
	//type myint int
	//var n1 int = 40
	//var n2 myint = 50
	//n1 =int(n2)

	res := funcpara2(getAdd, 500, 600)
	fmt.Println(res)

}
func getVal(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
func test08() {
	sum, sub := getVal(20, 10)
	fmt.Println(sum, sub)
}
func getAllsum(n1 int, anys ...int) int {
	sum := n1
	for _, val := range anys {
		sum += val
	}
	return sum

}
func test09() {
	res := getAllsum(10, 20, 30, 50, -20, -10, 70)
	fmt.Println(res)
}

var InitName string
var InitAge int
var InitTest = myInitTest()

func init() {
	InitName = "hello"
	InitAge = 99
	//fmt.Println(InitName, InitAge, "init在中间")
}
func myInitTest() int {
	//fmt.Println("myInitTest 变量先输出")
	return 90
}
func test10() {

	fmt.Println("mian最后:", myfunc.Name, myfunc.Age)
	fmt.Println("mian最后:", myfunc2.Name2, myfunc2.Age2)
}

var (
	b func(int, int) int = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func test11() {

	//一次性
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(res1)

	//可重复使用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 20)
	fmt.Println(res2)

	//全局匿名函数
	res3 := b(10, 20)
	fmt.Println(res3)
}
func AddUpper() func(int) int {
	n := 10
	return func(i int) int {
		n += i
		return n
	}
}
func test12() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
func makeSuffix(suffix string) func(string) string {
	return func(s string) string {
		if strings.HasSuffix(s, suffix) {
			return s
		} else {
			return s + suffix
		}
	}
}

func test13() {
	//写一个闭包，传入一个文件名进入闭包中，如果改文件名包含后缀，直接返回，不包含则加上后缀后返回
	f := makeSuffix(".jpg")
	fmt.Println(f("abc"))
	fmt.Println(f("a.jpg"))
}
func mydeferadd(n1 *int, n2 *int) int {
	fmt.Println("hahaha")
	return *n1 + *n2
}
func mydeferprint(n1 *int) {
	fmt.Println(*n1)
}
func deferfunc() {
	n1 := 10
	n2 := 20
	defer fmt.Println("ok1,res=", mydeferadd(&n1, &n2))
	defer mydeferprint(&n1)
	defer func() {
		fmt.Println("ok4,n1=", n1)
	}()
	defer func(n1 int) {
		fmt.Println("ok5,n1=", n1) //此时就是传参数了，而不是闭包类型了
	}(n1)
	defer fmt.Println("ok2,res=", mydeferadd(&n1, &n2))
	n1++
	n2++
	fmt.Println("ok3,res=", n1+n2)
}
func test14() {
	deferfunc()
	fmt.Println("ok4")
}
func f15test() (r int) {
	defer func() {
		r = r + 5
	}()
	return 1
}

func test15() {
	a := f15test()
	fmt.Println(a)

	for i := 0; i < 10; i++ {
		if i := 3; i < 5 {
			{
				i := 1
				fmt.Println(i)
			}
		}
	}

}
func test16() {

	num1 := 10
	fmt.Printf("numtype=%T,num1=%v,&num1=%v\n", num1, num1, &num1)
	num2 := new(int)
	fmt.Printf("num2type=%T,num2=%v,&num2=%v,*num2=%v\n", num2, num2, &num2, *num2)
	*num2 = 100
	fmt.Printf("num2type=%T,num2=%v,&num2=%v,*num2=%v\n", num2, num2, &num2, *num2)

}
func testerr(n1 float64, n2 float64) (res float64, err error) {
	if n2 == 0 {
		return -1, errors.New("除数不可以为0")
	}
	return n1 / n2, nil
}

func testerr2(n1 float64, n2 float64) (res float64) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if n2 == 0 {
		panic("除数不能为0")
	}
	return n1 / n2
}
func test17() {
	n1 := 10.0
	n2 := 0.0
	//return error
	res, err := testerr(n1, n2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("res is ", res)
	}

	//defer+panic+recover
	res2 := testerr2(n1, n2)
	fmt.Println(res2)
}
func funcB() {
	fmt.Println("func bbbb")
}
func funcC() {
	//str := recover()
	//fmt.Println(str)
	fmt.Println("func cccc")
}
func test18() {
	defer funcB()
	defer funcC()
	//fmt.Println("test18")
	panic("panicA")

}
func main() {
	//函数基本用法
	test01()

	//函数返回值
	//test02()

	//数组传参
	//test03()

	//地址传参
	//test04()

	//引用传参
	//test05()

	//函数类型
	//test06()

	//type类型定义
	//test07()

	//函数返回值定义
	//test08()

	//多个参数
	//test09()

	//init函数
	//test10()

	//匿名函数
	//test11()

	//闭包类型
	//test12()

	//闭包案例
	//test13()

	//defer
	//test14()

	//defer与return
	//test15()

	//new函数
	//test16()

	//错误机制
	//test17()

	//defer原理刨析
	//test18()
}
