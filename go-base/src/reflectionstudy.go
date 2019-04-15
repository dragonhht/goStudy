package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int8
}

func (user User) Study() {
	fmt.Println("%s study", user.Name)
}

func info(inter interface{}) {
	// 通过t可获取定义中的所有元素
	t := reflect.TypeOf(inter)
	fmt.Println("Type is ", t.Name())
	// 通过field可获取存储的值，也可去改变他
	filed := reflect.ValueOf(inter)
	// 获取所有的属性
	for i := 0; i < t.NumField(); i++ {
		// 获取属性和相应的值
		f := t.Field(i)
		val := filed.Field(i).Interface()
		fmt.Printf("%s : %v = %v\n", f.Name, f.Type, val)
	}

	// 获取所有的方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s : %v\n", m.Name, m.Type)
	}
}

func main() {
	u := User{"dragonhht", 18}
	info(u)
}
