package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
	"strconv"
	"regexp"
)

/*
	处理器.
 */
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析参数，默认不解析
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("schme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form{
		fmt.Println("key: " , k, ", val: ", v)
	}
	// 输出到客户端
	fmt.Fprint(w, "Hello World")
}

/*
	登录处理.
 */
func login(w http.ResponseWriter, r *http.Request) {
	// 获取请求方法
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		// 解析form
		r.ParseForm()
		// 执行登录逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		// 检验
		// 必填
		if len(r.Form.Get("username")) == 0 {
			fmt.Println("username为空")
		}

		// 转化为int数字
		getint, err := strconv.Atoi(r.Form.Get("age"))
		// 如果转换成功则为数字，否则则不是
		if err != nil {
			fmt.Println(err)
		}
		// 转换成功为数字
		fmt.Println("age: ", getint)
		// 使用正则表达式检验数字
		if m, _ := regexp.MatchString("[0-9]+$", r.Form.Get("age")); m {
			// 匹配成功
			fmt.Println("age is num :", m)
		}

		// 判断是否全为中文
		if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("username")); m {
			fmt.Println("username 为全中文")
		}
	}

}

func main() {
	// 设置访问的路由
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	// 设置监听端口
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
