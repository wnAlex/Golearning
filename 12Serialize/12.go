package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}
type Animal2 struct {
	Name string `json:"name"`
	//Age  int    `json:"age"`
	Sex  string `json:"sex"`
	Agg  int    `json:"agg"`
	Team string `json:"team"`
}

func test01() {
	ani := Animal{
		Name: "wnAlex",
		Age:  23,
		Sex:  "man",
	}
	aniMarshal, err := json.Marshal(ani)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(aniMarshal))

	//Unmarshal只针对字段，当反序列化对应的结构体有相应字段就成功对应，没有就对应失败取默认值，
	//如果结构体有多余字段也会对应失败取默认值
	ani2 := Animal2{}
	err = json.Unmarshal(aniMarshal, &ani2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ani2)
	fmt.Println("------------------------------")
}
func test02() {
	m := make(map[string]interface{}, 0)
	m["age"] = 12
	m["name"] = "wn"
	m["sex"] = "man!"

	marshalStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshalStr))

	//不需要重新分配内存，因为Unmarshal封装了make的操作
	var m2 map[string]interface{}

	//哪怕是引用类型，也需要传入地址
	err = json.Unmarshal(marshalStr, &m2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m2)
	fmt.Println("------------------------------")
}

func test03() {
	s := make([]map[string]interface{}, 0)
	m1 := make(map[string]interface{}, 0)
	m1["age"] = 12
	m1["name"] = "wn"
	m1["sex"] = "man!"
	s = append(s, m1)
	m2 := make(map[string]interface{}, 0)
	m2["age"] = 13
	m2["name"] = "molly"
	m2["sex"] = "woman!"
	s = append(s, m2)

	sliceMarshal, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sliceMarshal))

	var s2 []map[string]interface{}
	//不需要重新分配内存，因为Unmarshal封装了make的操作

	//哪怕是引用类型，也需要传入地址
	err = json.Unmarshal(sliceMarshal, &s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s2)
	fmt.Println("------------------------------")
}
func test04() {
	a := 10

	marshalStr, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshalStr))
}
func main() {
	//1.结构体序列化
	test01()

	//2.map序列化
	test02()

	//3.slice序列化
	test03()

	//4.int序列化
	test04()
}
