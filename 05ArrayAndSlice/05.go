package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

func test01() {
	//数组的定义
	//var arr3 [3]int
	//var arr4 [4]int = [4]int{1, 2, 3, 4}
	//var arr5 = [4]int{5, 6, 7, 8}
	//arr6 := [4]int{9, 10, 11, 12}
	//var arr7 [4]int = [4]int{1: 14, 0: 13, 3: 16, 2: 15}
	//var arr8 = [...]int{17, 18, 19, 20}
	//
	//fmt.Println(arr3)
	//fmt.Println(arr4)
	//fmt.Println(arr5)
	//fmt.Println(arr6)
	//fmt.Println(arr7)
	//fmt.Println(arr8)

	//数组可以直接打印,默认值是0
	//var arr1 [3]int
	//fmt.Println(arr1)
	//fmt.Printf("%v", arr1)

	//数组的首地址
	//arr2 := [3]int{1, 2, 3}
	//fmt.Printf("%p\n", &arr2)
	//fmt.Println(&arr2[0])
	//fmt.Println(&arr2[1])
	//fmt.Println(&arr2[2])

	//数组长度也是类型的一部分
	//arr10 := [3]int{1, 2, 3}
	//arr11 := [3]int{4, 5, 6}
	////arr11 := [4]int{4, 5, 6} 就没法赋值了
	//arr10 = arr11
	//
	//fmt.Println(arr10)
	//fmt.Println(arr11)

}
func test02() {
	arr1 := [3]string{"haha", "mama", "papa"}
	fmt.Println(arr1)
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for _, val := range arr1 {
		fmt.Println(val)
	}
}
func noChangeMyArr(arr [3]int) {
	for i, _ := range arr {
		arr[i]++
	}
}
func changeMyArr(arr *[3]int) {
	for i, _ := range arr {
		arr[i]++
	}
}
func test03() {
	arr1 := [3]int{1, 2, 3}
	for _, val := range arr1 {
		fmt.Println(val)
	}
	fmt.Println()
	noChangeMyArr(arr1)
	for _, val := range arr1 {
		fmt.Println(val)
	}
	fmt.Println()
	changeMyArr(&arr1)
	for _, val := range arr1 {
		fmt.Println(val)
	}
}

func test04() {
	rand.Seed(time.Now().Unix())
	var arr12 [5]int
	for i := 0; i < len(arr12); i++ {
		arr12[i] = rand.Intn(100)
		fmt.Printf("%d ", arr12[i])
	}
	fmt.Println()

	length := len(arr12)
	i := 0
	j := length - 1
	for i <= j {
		temp := arr12[j]
		arr12[j] = arr12[i]
		arr12[i] = temp
		i++
		j--
	}
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", arr12[i])
	}
	fmt.Println()

}
func test05() {
	arr1 := [...]int{1, 2, 3, 45, 65, 67, 87}
	arr2 := arr1[1:5]
	fmt.Println(arr2)

	fmt.Printf("arr2的地址:%p,arr2的指针指向的地址:%p,arr2的len:%d,arr2的cap:%d\n", &arr2, &arr2[0], len(arr2), cap(arr2))
	fmt.Printf("arr1的地址:%p,arr1的1号下标元素的地址:%p\n", &arr1, &arr1[1])

	arr2 = append(arr2, 100)
	arr2 = append(arr2, 200)
	arr2 = append(arr2, 300)
	arr2 = append(arr2, 400)
	arr2 = append(arr2, 500)
	//arr2[10] = 101000
	fmt.Println(arr2)
	fmt.Println(arr1)
}
func test06() {

	//1.直接定义
	var s1 []int
	fmt.Println(s1)
	var s2 []int = []int{1, 2, 3}
	fmt.Println(s2)

	//2. 省略类型
	var s3 = []int{4, 5, 6}
	fmt.Println(s3)

	//3. 自动类型推导
	s4 := []int{7, 8, 9}
	fmt.Println(s4)

	//4. 指定下标
	s5 := []int{1: 11, 0: 10, 2: 12, 4: 14, 3: 13}
	fmt.Println(s5)

	//5. 数组下标截取
	arr1 := [...]int{1, 2, 3, 45, 65, 67, 87}
	s6 := arr1[1:5]
	//s11 := s6[0:8]
	//fmt.Println(s11)
	fmt.Println(s6)

	//6. make函数分配
	s7 := make([]int, 4, 10)
	fmt.Println(s7)
	s7[1] = 100
	s7[3] = 300
	//s7[6] = 600 越界，index不能超过len
	fmt.Println(s7)

	s8 := make([]int, 0)
	fmt.Println(s8)
	//s8[0] = 1000 越界，只能用append方式追加内容
	fmt.Println(len(s8), cap(s8))

	s9 := make([]int, 2, 10)
	fmt.Println(s9)
	fmt.Println(len(s9), cap(s9))
	s10 := append(s9, 100, 200)
	fmt.Println(s9)
	fmt.Println(s10)
	fmt.Printf("s9地址:%p\n", &s9)
	fmt.Printf("s10地址:%p\n", &s10)
	fmt.Printf("s9对应元素地址:%p\n", &s9[0])
	fmt.Printf("s10对应元素地址:%p\n", &s10[0])
}

func test07() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 0, 10)
	fmt.Println(s1)
	fmt.Println(s2)
	num := copy(s2, s1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(num)
}

type stringStruct struct {
	str unsafe.Pointer //str首地址
	len int            //str长度
}
type sliceStruct struct {
	str unsafe.Pointer //slice首地址
	len int            //slice 长度
	cap int            //slice cap
}

func test08() {
	s := "hellohahaha"
	s1 := s[0:5]
	fmt.Println(s1)
	fmt.Printf("string的首元素地址:%p\n", (*stringStruct)(unsafe.Pointer(&s)).str)
	fmt.Printf("slice的首元素地址:%p\n", (*sliceStruct)(unsafe.Pointer(&s1)).str)

}

func test09() {
	//数组的定义
	//1.直接定义
	var arr3 [3][4]int
	var arr4 [3][4]int = [3][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	//2.初始化
	var arr5 = [3][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	//3.自动类型推导
	arr6 := [3][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	//4.指定下标
	var arr7 [3][4]int = [3][4]int{{0: 1, 2: 3}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	//5.自动推导行数、无法自动推导列数，列数必须确定，如果空着不写就直接变成slice了
	var arr8 = [...][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}
	//var arr9 = [3][...]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	fmt.Printf("arr3 type is %T\n", arr3)
	fmt.Println(arr3)
	fmt.Printf("arr4 type is %T\n", arr4)
	fmt.Println(arr4)
	fmt.Printf("arr5 type is %T\n", arr5)
	fmt.Println(arr5)
	fmt.Printf("arr6 type is %T\n", arr6)
	fmt.Println(arr6)
	fmt.Printf("arr7 type is %T\n", arr7)
	fmt.Println(arr7)
	fmt.Printf("arr8 type is %T\n", arr8)
	fmt.Println(arr8)
	//fmt.Printf("arr9 type is %T\n", arr9)
	//fmt.Println(arr9)
}
func test10() {
	var arr4 [3][4]int = [3][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	//1.直接遍历
	for i := 0; i < len(arr4); i++ {
		for j := 0; j < len(arr4[0]); j++ {
			fmt.Printf("%d ", arr4[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	//2.for range遍历
	for _, val := range arr4 {
		for _, val2 := range val {
			fmt.Printf("%d ", val2)
		}
		fmt.Println()
	}
}
func main() {
	//数组的使用
	//test01()

	//数组的遍历
	//test02()

	//数组传参
	//test03()

	//反转数组
	//test04()

	//slice的使用
	//test05()

	//slice的定义
	//test06()

	//slice深拷贝
	//test07()

	//slice与string
	//test08()

	// 二维数组的定义
	test09()

	//二维数组的遍历
	//test10()
}
