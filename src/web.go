package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
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
