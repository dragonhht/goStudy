package main

import (
	"fmt"
	"reflect"
	. "time"
)

// 变量申明
var count int = 32
var flo complex64
var a, b, c, d int = 1, 2, 3, 4

var (
	s1 int
	s2 uint
	s3 string
	s4 bool
)

// 常量申明
const testCount, d1, d3  = 33, "Hello", true
const d4 int = 32

const (
	q1 = "Hello"
	q2 = 32
	q3
)

const constCount  = iota

const constCount1  = iota

const (
	constCount2 = iota
	_	// 跳值使用
	constCount3 = iota
)

func print() string  {
	return "World"
}

/*
	if-else测试.
 */
func ifTest() {
	a := 2
	if a > 0 {
		fmt.Println("a > 0")
	} else {
		if a == 0 {
			fmt.Println("a == 0")
		} else {
			fmt.Println(" a <  0")
		}
	}
}

func switchTest(num int) {
	switch num {
	case 1:
		fmt.Println("this is 1")
	case 2:
		fmt.Println("this is 2")
	default:
		fmt.Println("this is not 1 or 2")
	}
}

/**
	无限循环.
 */
func forTest1() {
	count := 1
	for {
		count++
		fmt.Println(count)
		Sleep(1 * Second)
	}
}

func forTest2() {
	for i := 0; i < 5; i++{
		fmt.Println(i)
	}
}

/**
	foreach。
 */
func forTest3() {
	arr := []string{"a", "b", "c", "d", "e"}
	for key, value := range arr {
		fmt.Printf("%d and %s", key, value)
	}
}

/**
	while
 */
func forTest4() {
	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}
}

/**
	goto
 */
func gotoTest() {
	goto One
	fmt.Println("this after One")
	One:
		fmt.Println("after")
}

func main() {
	n1, n2 := 2, 3
	fmt.Print(n1 + n2)
	fmt.Println(reflect.TypeOf(a))
	var z = int64(a)
	fmt.Println(reflect.TypeOf(z))
	fmt.Println(Time{})
	fmt.Println(&n1)
	fmt.Println(testCount)
	fmt.Println(len(q1))
	fmt.Println(print())
	fmt.Println(constCount)
	fmt.Println(constCount3)
	fmt.Println(q3)
	ifTest()
	switchTest(3)
	forTest2()
	forTest3()
	forTest4()
	gotoTest()
}
