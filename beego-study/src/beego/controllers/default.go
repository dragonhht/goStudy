package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user.html"
}

type UserController struct {
	beego.Controller
}

type User struct {
	Id int
	Username interface{} `form:"username"`
	Email string
}

// @router /user/get/:id [get]
func (u *UserController) Get() {
	//id := u.Ctx.Input.Param(":id")
	id, err := u.GetInt(":id")
	if err == nil {
		log.Printf("id is %v", id)
	}
	u.Data["message"] = "hello world"
	log.Println("UserController GEt")
	u.TplName = "index.html"
}

// @router /user [post]
func (this *UserController) Post() {
	// 获取上传的文件
	f, h, err := this.GetFile("uploadFile")
	log.Printf("文件Size: %v", h.Size)
	if err != nil {
		log.Println(err)
	}
	// 将上传的文件保存
	defer f.Close()
	this.SaveToFile("uploadFile", "static/upload/" + h.Filename)

	user := User{}
	if err := this.ParseForm(&user); err != nil {
		log.Println(err)
	}
	log.Println(user)
	this.Ctx.Redirect(302, "/user/get/12")
}
