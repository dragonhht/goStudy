package main

import (
	"fmt"
	"reflect"
	. "time"
)

var count int = 32
var flo complex64

const testCount, d1, d3  = 33, "Hello", true

var a, b, c, d int = 1, 2, 3, 4

var (
	s1 int
	s2 uint
	s3 string
	s4 bool
)

func main() {
	n1, n2 := 2, 3
	fmt.Print(n1 + n2)
	fmt.Println(reflect.TypeOf(a))
	var z = int64(a)
	fmt.Println(reflect.TypeOf(z))
	fmt.Println(Time{})
	fmt.Println(&n1)
	fmt.Println(testCount)
}
