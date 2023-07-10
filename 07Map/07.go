package main

import (
	"fmt"
	"sort"
)

func test01() {
	//1.直接定义
	var m map[string]string
	fmt.Println("m:", m)
	//m["1"] = "a" 没有分配内存，没法直接使用
	var m1 map[string]string = map[string]string{"1": "a", "2": "b"}
	fmt.Println("m1:", m1)
	m1["1"] = "c"
	fmt.Println("m1:", m1)

	//2.省略类型定义
	var m3 = map[string]string{"1": "a", "2": "b"}
	fmt.Println("m3:", m3)

	//3.自动类型推导
	m4 := map[string]string{"1": "a", "2": "b"}
	m4["3"] = "c"
	fmt.Println("m4:", m4)

	//4.make分配
	m5 := make(map[string]string)
	fmt.Println("m5 len:", len(m5))
	fmt.Println("m5:", m5)
	m5["1"] = "a"
	m5["2"] = "b"
	fmt.Println("m5:", m5)
	m5["3"] = "c"
	fmt.Println("m5:", m5)
	fmt.Println("m5 len:", len(m5))
}
func test02() {
	//1.map的key第一次出现就是增加，后续出现就是改
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "b"
	fmt.Println(m)
	m[1] = "c"
	fmt.Println(m)

	//2.map的删除
	delete(m, 1)
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)

	// 删除所有
	m = make(map[int]string)
	fmt.Println(m)

	//3.map的查找
	m[1] = "a"
	m[2] = "b"
	value, isHave := m[1]
	fmt.Printf("vlaue :%T,isHave :%T\n", value, isHave)
	fmt.Println(value, " ", isHave)
}
func test03() {
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "b"
	m[3] = "a"
	m[4] = "b"
	fmt.Println(m)
	for key, val := range m {
		fmt.Println(key, ":", val)
	}
}
func test04() {
	var ms []map[string]string = make([]map[string]string, 0)
	m1 := map[string]string{
		"name": "aa",
		"sex":  "man",
	}
	ms = append(ms, m1)
	m2 := map[string]string{
		"name": "bb",
		"sex":  "woman",
	}
	ms = append(ms, m2)
	fmt.Println(ms)
}
func test05() {
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "b"
	m[3] = "a"
	m[4] = "b"
	s := make([]int, 0)
	for key, val := range m {
		fmt.Println(key, ":", val)
		s = append(s, key)
	}
	sort.Ints(s)
	fmt.Println("After Sort")
	for _, val := range s {
		fmt.Printf("%d : %s\n", val, m[val])
	}
}
func main() {
	//map的定义
	//test01()

	//map的增删改查
	//test02()

	//map的遍历
	//test03()

	//map的slice
	//test04()

	//map的排序
	test05()
}
