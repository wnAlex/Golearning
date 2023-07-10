package main

import "fmt"

func BubbleSort(s []int) {
	for j := len(s) - 1; j > 0; j-- {
		flag := true
		for i := 0; i < j; i++ {
			if s[i] > s[i+1] {
				temp := s[i+1]
				s[i+1] = s[i]
				s[i] = temp
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
func test01() {
	s1 := []int{12, 4, 56, 71, 2, 5}
	fmt.Println(s1)
	BubbleSort(s1)
	fmt.Println(s1)
}
func binarySearch1(s []int, left int, right int, target int) int {
	if left == right {
		if s[left] == target {
			return left
		} else {
			return -1
		}
	}
	mid := (left + right) / 2
	if s[mid] == target {
		return mid
	} else if s[mid] > target {
		return binarySearch1(s, left, mid-1, target)
	} else {
		return binarySearch1(s, mid+1, right, target)
	}
}
func binarySearch2(s []int, target int) int {
	left := 0
	right := len(s) - 1
	for left <= right {
		mid := (left + right) / 2
		if s[mid] > target {
			right = mid - 1
		} else if s[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
func test02() {
	//1.递归实现
	s1 := []int{12, 4, 56, 71, 2, 5}
	BubbleSort(s1)
	fmt.Println(s1)
	target := 5
	index := binarySearch1(s1, 0, len(s1)-1, target)
	fmt.Println(index)
	//2.循环实现
	index2 := binarySearch2(s1, target)
	fmt.Println(index2)
}
func main() {
	//bubbleSort
	//test01()

	//二分查找
	//test02()
}
