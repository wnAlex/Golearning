package main

import (
	"fmt"
	"strconv"
	"strings"
)

func test06() {
	var s1 float64 = 2
	fmt.Println(s1)

	//错误使用
	//a:=10
	//b:=a++
}
func test07() {
	//1 Scanln 当出现类型不一致的时候会默认赋值失败
	var age int
	var name string
	var isWorked bool
	var salary float64

	//2 Scanln split切分
	fmt.Println("请输入姓名，年龄，是否工作，薪水：")
	var str string
	_, err := fmt.Scanln(&str)
	if err != nil {
		fmt.Println(err)
	}
	message := strings.Split(str, ",")
	name = message[0]
	age, err = strconv.Atoi(message[1])
	if err != nil {
		fmt.Println(err)
	}
	isWorked, err = strconv.ParseBool(message[2])
	if err != nil {
		fmt.Println(err)
	}
	salary, err = strconv.ParseFloat(message[3], 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("name = %s,age=%d,isWorked=%t,salary=%f\n", name, age, isWorked, salary)

	//3 Scanf
	fmt.Println("请输入姓名，年龄，是否工作，薪水：")
	_, err = fmt.Scanf("%s,%d,%t,%f", &name, &age, &isWorked, &salary)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("name = %s,age=%d,isWorked=%t,salary=%f\n", name, age, isWorked, salary)

}
func test08() {

	var i int = 10
	//10 to 2
	temp2 := strconv.FormatInt(int64(i), 2)
	i2, _ := strconv.ParseInt(temp2, 10, 0)
	fmt.Println(i2)

	//10 to 8
	temp3 := strconv.FormatInt(int64(i), 8)
	i3, _ := strconv.ParseInt(temp3, 10, 0)
	fmt.Println(i3)

	//10 to 16 只能变成字符串
	temp4 := strconv.FormatInt(int64(i), 16)
	fmt.Println(temp4)
	i4, _ := strconv.ParseInt(temp4, 16, 0)
	fmt.Println(i4)

	//2 to 10  Go中无法用二进制直接表示一个数
	temp2 = strconv.FormatInt(int64(i), 2)
	i2, _ = strconv.ParseInt(temp2, 10, 0)

	//8 to 10
	i = 011
	temp5 := strconv.FormatInt(int64(i), 10)
	fmt.Println(temp5)
	i5, _ := strconv.ParseInt(temp5, 10, 0)
	fmt.Println(i5)

	//16 to 10
	i = 0x11
	temp6 := strconv.FormatInt(int64(i), 16)
	fmt.Println(temp6)
	i6, _ := strconv.ParseInt(temp6, 10, 0)
	fmt.Println(i6)

}
func main() {
	//运算符
	//test06()

	//读取输入
	//test07()

	//进制
	//test08()
}
