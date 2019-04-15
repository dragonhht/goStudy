package main

import "fmt"

/*
	定义结构体.
 */
type Student struct {
	name string
	age int8
}

/*
	值传递，不改变原有的值
 */
func setStudentName(student Student) {
	student.name = "hello"
}

/*
	通过指针，改变了原有的值.
 */
func setStudentName1(student *Student) {
	student.name = "hello"
}

func main() {
	stu := Student{"dragonhht", 18}
	stu.age = 20
	fmt.Println(stu)
	setStudentName(stu)
	fmt.Println(stu)
	setStudentName1(&stu)
	fmt.Println(stu)
}
