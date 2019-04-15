package main

import "fmt"

func arrTest(arr []int) {
	arr[0] = 3
}

func main() {
	var arr [5] int
	arr[0] = 1
	fmt.Println(arr)

	// 方式一, 指定数组大小
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	// 方式二, 不手动设置数组大小，通过元素个数自动设置数组大小
	var arr2 = []int{1, 2, 4, 5}
	fmt.Println(arr2)
	fmt.Println(arr2[2])
	arrTest(arr2)
	fmt.Println(arr2)
}
