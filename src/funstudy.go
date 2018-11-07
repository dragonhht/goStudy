package main

import "fmt"

func fun1(n int) {
	n = 9
}

func fun2(n *int) {
	*n = 9
}

func fun3(n1, n2 int) int {
	return n1 + n2
}

func fun4(str1, str2 string) (string, string) {
	return str1, str2
}

func fun5(str1 []string) {
	str1[0] = "hello"
}

/*
	闭包.
 */
func fun7(n1, n2 int) func(n3, n4 int) int {
	i := n1
	return func(n3, n4 int) int {
		return i + n2 + n3 + n4
	}
}

/*
	定义结构体.
 */
type type1 struct {
	name string
}

/*
	定义属于type1类型的方法.
 */
func (t type1) getName() string {
	return t.name
}

func main() {
	a := 4
	fun1(a)
	fmt.Println(a)
	fun2(&a)
	fmt.Println(a)
	fmt.Println(fun3(1 , 2))
	str1, str2 := fun4("Hello", "World")
	fmt.Printf("%s and %s\n", str1, str2)
	arr := []string{"a", "b"}
	fmt.Println(arr)
	fun5(arr)
	fmt.Println(arr)

	fun6 := func(n1, n2 int) int {
		return n1 + n2
	}
	fmt.Println(fun6(3, 4))

	fun7 := fun7(2, 3)
	fmt.Println(fun7(3, 4))
	fmt.Println(fun7(5, 6))

	var t type1
	t.name = "dragonhht"
	fmt.Println(t.getName())
}
