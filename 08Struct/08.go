package main

import (
	"encoding/json"
	"fmt"
	"github.com/wnAlex/Golearning/08Struct/student"
	"unsafe"
)

type Cat struct {
	Name string
	Sex  string
	age  int
	Paws byte
}

func test01() {

	var cat Cat
	cat.Sex = "woman"
	cat.Name = "molly"
	cat.Paws = 4
	cat.age = 1
	fmt.Printf("cat 的首地址:%p\n", &cat)
	fmt.Printf("cat 首元素的地址:%p\n", &cat.Name)
	fmt.Println("cat 的大小：", unsafe.Sizeof(cat))
	fmt.Println("cat name 的大小：", unsafe.Sizeof(cat.Name))
	fmt.Println("cat sex的大小：", unsafe.Sizeof(cat.Sex))
	fmt.Println("cat age的大小：", unsafe.Sizeof(cat.age))
	fmt.Println("cat paws的大小：", unsafe.Sizeof(cat.Paws))

}
func test02() {
	//1.直接定义
	var cat Cat
	fmt.Println(cat)
	var cat1 Cat = Cat{
		Name: "aa",
		Sex:  "woman",
		age:  13,
		Paws: 4,
	}
	fmt.Println(cat1)
	fmt.Println()

	//2.初始化
	var cat2 = Cat{
		Name: "bb",
		Sex:  "woman",
		age:  14,
		Paws: 4,
	}
	fmt.Println(cat2)
	fmt.Println()

	//3.自动类型推导
	cat3 := Cat{
		Name: "cc",
		Sex:  "man",
		age:  15,
		Paws: 4,
	}
	fmt.Println(cat3)
	fmt.Println()

	//4.new的方式
	cat4 := new(Cat)
	fmt.Println(cat4)
	cat4.Name = "dd"
	cat4.Sex = "woman"
	cat4.age = 16
	cat4.Paws = 4
	fmt.Println(cat4)
	fmt.Println()

	//5.直接匿名对象取地址
	var cat5 *Cat = &Cat{
		Name: "ee",
		Sex:  "woman",
		age:  17,
		Paws: 4,
	}
	fmt.Println(cat5)
	cat5.Sex = "man"
	fmt.Println(cat5)
}

type A struct {
	Num1 string
}
type B struct {
	Num2 string
}

func test03() {
	var a A
	var b B
	//a = A(b)
	a = *(*A)(unsafe.Pointer(&b))
	fmt.Println(a)
	fmt.Println(b)
}

type Monster struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func test04() {
	m := Monster{
		Name: "molly",
		Age:  25,
		Sex:  "woman",
	}
	jsonm, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonm))
}
func (m *Monster) modMonster(name string) {
	//fmt.Println("modMonster ", m.Name)
	m.Name = name
	fmt.Println("modMonster ", m.Name)
}
func (m Monster) modMonster2(name string) {
	//fmt.Println("modMonster ", m.Name)
	m.Name = name
	fmt.Println("modMonster ", m.Name)
}
func test05() {
	m := Monster{
		Name: "molly",
		Age:  25,
		Sex:  "woman",
	}
	m.modMonster("alex")
	fmt.Println("main ", m.Name)
	(&m).modMonster("blex")
	fmt.Println("main ", m.Name)
	m.modMonster2("clex")
	fmt.Println("main ", m.Name)
	(&m).modMonster2("dlex")
	fmt.Println("main ", m.Name)
}

func (m *Monster) String() string {
	str := fmt.Sprintf("monster name=%s,age=%d,sex=%s ", m.Name, m.Age, m.Sex)
	return str
}
func test06() {
	m := Monster{
		Name: "molly",
		Age:  25,
		Sex:  "woman",
	}
	fmt.Println(m)
	fmt.Println(&m)
}
func test07() {
	//1.如果student包中Student首字母大写，那么可以直接访问，如果小写则需要工厂模式
	s := student.GetStudent("molly", 18)
	fmt.Println(s)
	fmt.Println(s.Name)
	fmt.Println(s.Age)
	fmt.Println("---------")
	//2.指针模式
	s2 := student.GetStudentPointer("alex", 19)
	fmt.Println(s2)
	fmt.Println(s2.Name)
	fmt.Println(s2.Age)
	fmt.Println("---------")
	//3.如果字段也小写，那么也需要额外封装对应的get和set方法
	fmt.Println(s.GetSex())
	s.SetSex("man")
	fmt.Println(s.GetSex())
	fmt.Println(s2.GetSex())
	s2.SetSex("woman")
	fmt.Println(s2.GetSex())

}
func main() {
	//1.结构体是值类型
	//test01()

	//2.结构体的定义
	//test02()

	//3.结构体的转换
	//test03()

	//4.结构体序列化和反序列化--tag
	//test04()

	//5.方法
	//test05()

	//6.println与string()
	//test06()

	//工厂模式
	//test07()
}
