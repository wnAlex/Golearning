package main

import (
	"fmt"
	"github.com/wnAlex/Golearning/04functionAndPackage/myfunc"
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

func main() {
	//函数基本用法
	test01()

}
