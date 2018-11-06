package test1

import (
	"fmt"
	"test2"
)

func init()  {
	fmt.Println("this is test 1")
}

func Test() {
	test2.Test2()
}
