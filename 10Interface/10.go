package main

import (
	"fmt"
	"math/rand"
	"sort"
	"unsafe"
)

type Animal interface {
	sleep()
	work()
}
type Animal2 interface {
	sleep()
	work()
	play()
}
type Cat struct {
	Name string
	Age  int
}

func (c *Cat) sleep() {
	fmt.Println("cat is sleeping")
}
func (c *Cat) work() {
	fmt.Println("cat is working")
}

type Dog struct {
}

func (d *Dog) sleep() {
	fmt.Println("dog is sleeping")
}
func (d *Dog) work() {
	fmt.Println("dog is working")
}

type Person struct {
}

func (p *Person) startWork(a Animal) {
	a.work()
	a.sleep()
}
func test01() {
	p := Person{}
	c := Cat{
		Name: "molly",
		Age:  12,
	}
	c.sleep()
	p.startWork(&c)
	p.startWork(&Dog{})
}
func test02() {
	var a Animal
	fmt.Printf("%v,%T\n", a, a)
	var c *Cat
	fmt.Printf("%v,%T\n", c, c)

}

type Bnimaml interface {
	Animal
	work2()
	sleep2()
}
type Cog struct {
	s string
}

func (c Cog) sleep() {

}
func (c Cog) work() {

}
func (c Cog) sleep2() {
	fmt.Println("sleep2")
}
func (c Cog) work2() {

}
func test03() {
	var b Bnimaml
	b = &Cog{
		s: "aaa",
	}
	fmt.Printf("%v,%T\n", b, b)
	fmt.Println(b)

	b.sleep2()
}

type T interface {
}

func test04() {
	var t T
	n1 := 10
	n2 := 9.9
	fmt.Printf("%v,%T\n", t, t)
	t = n1
	fmt.Printf("%v,%T\n", t, t)
	t = n2
	fmt.Printf("%v,%T\n", t, t)

	var t2 interface{}
	fmt.Printf("%v,%T\n", t2, t2)
	t2 = n1
	fmt.Printf("%v,%T\n", t2, t2)
	t2 = n2
	fmt.Printf("%v,%T\n", t2, t2)
	fmt.Printf("%d\n", unsafe.Sizeof(t2))
	var a Animal
	var b Bnimaml
	fmt.Printf("%d\n", unsafe.Sizeof(a))
	fmt.Printf("%d\n", unsafe.Sizeof(b))

}

type Student struct {
	Name  string
	Grade int
	Age   int
}
type Stu []Student

func (s Stu) Len() int {
	return len(s)
}
func (s Stu) Less(i, j int) bool {
	if s[i].Grade == s[j].Grade {
		return s[i].Age < s[j].Age
	}
	return s[i].Grade < s[j].Grade
}
func (s Stu) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func test05() {
	s := Stu{}
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  "王楠" + fmt.Sprintf("%d", i),
			Grade: rand.Intn(10),
			Age:   rand.Intn(20),
		}
		s = append(s, stu)
	}
	for _, val := range s {
		fmt.Println(val)
	}
	sort.Sort(s)
	fmt.Println("aaaaaaaaaaaaaaaa-------------------")
	for _, val := range s {
		fmt.Println(val)
	}

}

func test06() {
	var t2 interface{}
	n1 := 10
	n2 := 9.9
	fmt.Printf("%v,%T\n", t2, t2)
	t2 = n1
	fmt.Printf("%v,%T\n", t2, t2)
	t2 = n2
	fmt.Printf("%v,%T\n", t2, t2)
	fmt.Println("-------------")
	n3 := t2
	fmt.Printf("%v,%T\n", n3, n3)
	var n4 float64
	//n4=t2
	//n4=n3
	//n4=float64(t2)
	n4 = t2.(float64)
	fmt.Printf("%v,%T\n", n4, n4)
	fmt.Println("-------------")
	n5, isOk := t2.(int)
	if !isOk {
		fmt.Println("fail")
	} else {
		fmt.Println(n5)
	}
	fmt.Println("-------------")

}

func main() {
	//1.接口
	//test01()

	//2.接口的类型
	//test02()

	//3.接口的继承
	test03()

	//4.空接口
	//test04()

	//5.实现sort接口
	//test05()

	//6.Assertion
	//test06()

}
