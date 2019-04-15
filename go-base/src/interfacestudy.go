package main

import "fmt"

/*
	定义接口.
 */
type Interface1 interface {
	getName() string
	setName(name string)
}

/*
	定义结构体.
 */
type Struct1 struct {
	name string
}

/*
	实现接口.
 */
func (struct1 Struct1) getName() string  {
	return struct1.name
}

/*
	实现接口.
 */
func (struct1 *Struct1) setName(name string) {
	struct1.name = name
}

/*
	实现错误接口.
 */
func (struct1 *Struct1) Error() string {
	msg := "错误信息"
	return fmt.Sprintln(msg, struct1.name)
}

func main() {
	stu := Struct1{"dragonhht"}
	fmt.Println(stu)
	fmt.Println(stu.getName())
	stu.setName("huang")
	fmt.Println(stu)
	fmt.Println(stu.getName())
}
