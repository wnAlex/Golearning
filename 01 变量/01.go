package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
	"unsafe"
)

func test01() {
	var a1 int
	a1 = 10
	fmt.Println(a1)

	var a2 = 20
	fmt.Println(a2)

	a3 := 30
	fmt.Println(a3)

	var a4, a5 int
	a4 = 40
	a5 = 50
	fmt.Println(a4, a5)

	var a6, a7 = 60, 70
	fmt.Println(a6, a7)

	a8, a9 := 80, "aaa"
	fmt.Println(a8, a9)

	var (
		a10 = 100.1
		a11 = "bbb"
	)
	fmt.Println(a10, a11)
	fmt.Printf("%T %T %T\n", a8, a9, a10)
	fmt.Printf("a9 type= %T,a10 size = %d", a9, unsafe.Sizeof(a10))
}
func test02() {
	a8, a9 := 80, "aaa"
	fmt.Println(a8, a9)

	var (
		a10 = 100.1
		a11 = "bbb"
	)
	fmt.Println(a10, a11)
	fmt.Printf("%T %T %T\n", a8, a9, a10)
	fmt.Printf("a9 type= %T,a10 size = %d", a9, unsafe.Sizeof(a10))
}
func test03() {
	var a byte = 'a'
	fmt.Println(a)

	var b = 'b'
	fmt.Printf("b=%c\n", b)

	//var c byte='王'
	//fmt.Println(c)

	c := '人'       //和默认一样都是rune类型接受汉字字符，采用unicode编码
	fmt.Println(c) //输出的是unicode编码
	fmt.Printf("c=%d,c= %c,c type =%T\n", c, c, c)
	fmt.Println("len(\"c\"):", unsafe.Sizeof(c))

	var d rune = rune('人') //和默认一样都是rune类型接受汉字字符，采用unicode编码
	fmt.Println(d)         //输出的是unicode编码
	fmt.Printf("d=%d,d= %c,d type =%T\n", d, d, d)
	fmt.Println("len(\"d\"):", unsafe.Sizeof(d))

	var e int = 20154
	fmt.Printf("e= %c,e type =%T\n", e, e)

	bytestr := "try学习"
	fmt.Println("len(\"try学习\"):", len(bytestr))                                       //9
	fmt.Println("utf8.RuneCountInString(\"try学习\"):", utf8.RuneCountInString(bytestr)) //5
	fmt.Println("len([]rune(try学习)):", len([]rune("try学习")))
	fmt.Println("sizeof(try学习):", unsafe.Sizeof(bytestr))                 //返回的是字符串类型的大小 为16
	fmt.Println("sizeof([]rune(try学习)):", unsafe.Sizeof([]rune(bytestr))) //返回的是int32数组的大小 24
	fmt.Printf("rune(hex):%x\n", []rune(bytestr))
	fmt.Printf("bytes(hex):% x\n", []byte(bytestr))
}

type stringStruct struct {
	str unsafe.Pointer //str首地址
	len int            //str长度
}

func test04() {
	var str string = "aa\\n"
	var str2 string = `aa\\n`
	fmt.Println(str)
	fmt.Println(str2)

	str3 := "asd" + "asd "
	fmt.Println(str3)
	for index, c := range str3 {
		fmt.Printf("%d %c\n", index, c)
		c = 'm'
		fmt.Printf("%c\n", c)
	}
	fmt.Println(str3)

	str4 := "eeeee"
	fmt.Println(str4)
	for index, c := range str4 {
		fmt.Printf("%d %c\n", index, c)
		c = 'm'
		fmt.Printf("%c\n", c)
	}
	fmt.Println(str4)

	str5 := "qwe" + "rty "
	str6 := "yui"
	i := 10
	var mystr stringStruct = *(*stringStruct)(unsafe.Pointer(&str5))
	var mystr2 stringStruct = *(*stringStruct)(unsafe.Pointer(&str6))
	fmt.Println(mystr.str)
	fmt.Println(mystr.len)
	fmt.Println(mystr2.str)
	fmt.Println(mystr2.len)
	fmt.Println(&i)

	s := false
	fmt.Printf("%v", s)

}
func test05() {
	i := 10
	var f float32 = float32(i)
	fmt.Println(f)

	var s1 int = 10
	var s2 float64 = 20.01
	var s3 bool = false

	str0 := strconv.Itoa(s1)
	fmt.Printf("str0 type =%T,value=%s\n", str0, str0)

	str1 := fmt.Sprintf("%v+%v+%v", s1, s2, s3)
	fmt.Printf("str1 type =%T,value=%s\n", str1, str1)

	str2 := strconv.FormatInt(int64(s1), 2)
	fmt.Printf("str2 type =%T,value=%s\n", str2, str2)

	str3 := strconv.FormatBool(s3)
	fmt.Printf("str3 type =%T,value=%s\n", str3, str3)

	str4 := strconv.FormatFloat(s2, 'f', 10, 64)
	fmt.Printf("str4 type =%T,value=%s\n", str4, str4)

	str5, _ := strconv.Atoi("10")
	fmt.Printf("str5 type =%T,value=%d\n", str5, str5)

	//返回类型是int64，bitsize存在的作用就是在转换过程中做了限制，如果是8，那么转化最大值就是127，然后把这个结果放在int64中返回。
	str6, err := strconv.ParseInt("1011111111111111", 10, 8)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("str6 type =%T,value=%d\n", str6, str6)

	//返回类型是float64，bitsize参数的作用就是在转换时做了限制，如果是32，那么就只能保证7位精度，然后再将这个结果放到float64中
	str7, err := strconv.ParseFloat("230.2012223312123", 32)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("str7 type =%T,value=%v\n", str7, str7)
}

const (
	BEIJING  = 1
	SHANGHAI = 2
	SHENZHEN = 3
)
const (
	a, b = iota * 10, iota*10 + 1
	c, d
	e, f
	g, h = iota * 3, iota*3 + 1
	i, j
)

func test06() {
	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
	fmt.Println(SHENZHEN)
	fmt.Println("---------")
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
	fmt.Println(g, h)
	fmt.Println(i, j)
}
func main() {
	//变量的定义
	//test01()

	//整形和浮点型的使用
	//test02()

	//字符的使用
	//test03()

	//字符串的使用
	//test04()

	//基本数据类型转换
	//test05()

	//const iota
	test06()
}
