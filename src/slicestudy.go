package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	//var slice1 []int = make([]int, 5)
	//slice2 := [] int {1, 2, 3}
	slice3 := arr[:]
	fmt.Println(slice3[1:4])
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 55);fmt.Println(slice3)
	var slice4 = make([]int, 5, 6)
	copy(slice4, slice3)
	fmt.Println(slice4)
}
