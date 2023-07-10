package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) GetName() string {
	return a.Name
}
func (a *Animal) SetName(name string) {
	a.Name = name
}
func (a *Animal) GetAge() int {
	return a.Age
}
func (a *Animal) SetAge(age int) {
	a.Age = age
}

type Cat struct {
	Animal
	Paws int
}

func (c *Cat) Running() {
	fmt.Println("cat is running")
}

type Bird struct {
	Wings int
	Animal
}

func (b *Bird) Flying() {
	fmt.Println("bird is flying")
}

func test01() {
	c := Cat{
		Animal: Animal{
			Name: "kitty",
			Age:  1,
		},
		Paws: 4,
	}
	fmt.Println(c.GetAge())
	c.Running()
	fmt.Printf("c.Name=%p,c.Age=%p,c.Paws=%p\n", &c.Name, &c.Age, &c.Paws)
	fmt.Println()
	b := Bird{
		Wings: 2,
		Animal: Animal{
			Name: "alex",
			Age:  2,
		},
	}
	fmt.Println(b.GetName())
	b.Flying()
	fmt.Printf("b.Name=%p,b.Age=%p,b.Wings=%p\n", &b.Name, &b.Age, &b.Wings)
}

type Base struct {
	Name string
	Age  int
}

func (b *Base) test01() {
	fmt.Println("base::", b.Name)
}
func (b *Base) test02() {
	fmt.Println("base::", b.Name)
}

type Son struct {
	Base
	Name string
}

func (s *Son) test02() {
	fmt.Println("son::", s.Name)
}

func test02() {
	s := Son{}
	s.Name = "molly"
	fmt.Println("s.name:", s.Name)
	fmt.Println("s.base.name:", s.Base.Name)
	s.test02()
	s.test01()
	s.Base.test02()
}

type Base1 struct {
	Name string
	Age  int
}
type Base2 struct {
	Name string
	Age  int
}
type Son1 struct {
	*Base1
	*Base2
	Name string
}

func test03() {
	s := Son1{
		Base1: &Base1{
			Name: "Base1",
			Age:  10,
		},
		Base2: &Base2{
			Name: "Base2",
			Age:  20,
		},
		Name: "Son1",
	}
	fmt.Println(s)
	fmt.Println(*s.Base1, *s.Base2, s.Name)
}

type Base3 struct {
	Name string
	Age  int
}
type Son3 struct {
	Base3
	int
	n int
}

func test04() {
	s := Son3{
		Base3: Base3{
			Name: "Alex",
			Age:  10,
		},
		int: 0,
		n:   10,
	}
	fmt.Println(s)
	s.int = 20
	s.n = 30
	fmt.Println(s)
}
func main() {
	//1.继承
	//test01()

	//2.出现同名变量
	//test02()

	//3.继承指针
	//test03()

	//4.继承内置数据类型
	test04()
}
