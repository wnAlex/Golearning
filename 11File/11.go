package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func test01() {
	file, err := os.Open("H:\\Go project\\src\\Golearning\\11File\\a.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件打开成功")
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取完毕")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件关闭成功")

}
func test02() {
	filename := "H:\\Go project\\src\\Golearning\\11File\\a.txt"

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件读取成功")
	fmt.Println(string(content))
}

func test03() {

	file, err := os.OpenFile("H:\\Go project\\src\\Golearning\\11File\\a.txt", os.O_CREATE|os.O_APPEND, fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err = writer.WriteString("hahaha\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	//刷新缓冲
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件关闭成功")

}
func test04() {
	file, err := os.Open("H:\\Go project\\src\\Golearning\\11File\\a.txt")
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
		fmt.Println(str)
	}

	// open打开的文件默认只读
	//writer := bufio.NewWriter(file)
	//str := "hahahahannnnnnnnn\r\n"
	//for i := 0; i < 10; i++ {
	//	_, err = writer.WriteString(str)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件关闭成功")
}
func test05() {
	filepath := "H:\\Go project\\src\\Golearning\\11File\\b.txt"
	fileinfo, err := os.Stat(filepath)
	if err == nil {
		fmt.Println("file exist")
	} else if os.IsNotExist(err) {
		fmt.Println("file not exist")
		return
	}
	fmt.Println(fileinfo.Name())

}
func test06() {
	filepath := "H:\\Go project\\src\\Golearning\\11File\\xx.png"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件打开成功")
	reader := bufio.NewReader(file)
	p := make([]byte, 1000000)
	for {
		_, err := reader.Read(p)
		if err == io.EOF {
			break
		}

	}
	fmt.Printf("%v", p)
	fmt.Println("文件读取完毕")

}

func test07() {
	filepath := "H:\\Go project\\src\\Golearning\\11File\\xx.png"
	readFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(readFile)

	filepath2 := "H:\\Go project\\src\\Golearning\\asd\\xx.png"
	writeFile, err := os.OpenFile(filepath2, os.O_CREATE|os.O_WRONLY, 0)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(writeFile)

	n, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("拷贝完成,拷贝了", n, "字节")

}
func test08() {
	var userName string
	var userAge int
	//注册解析规则
	flag.StringVar(&userName, "u", "默认值", "user的name")
	flag.IntVar(&userAge, "a", 0, "user的age")

	//开始解析
	flag.Parse()

	fmt.Println(userName)
	fmt.Println(userAge)
}
func main() {
	//1.带缓冲区的读取文件
	//test01()

	//2.一次性读取文件
	//test02()

	//3.创建文件并用bufio写入
	//test03()

	//4.读取文件并追加内容
	//test04()

	//5.判断文件是否存在
	//test05()

	//6.读取二进制文件
	//test06()

	//7.拷贝文件
	//test07()

	//8.解析命令行参数
	test08()

}
