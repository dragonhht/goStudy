package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type SessionController struct {
	beego.Controller
}

// @router /session [get]
func (this *SessionController) Get()  {
	// 获取Session中的值
	username := this.GetSession("userName")
	log.Printf("userName is %v", username)
	this.TplName = "index.html"
}

type Session2Controller struct {
	beego.Controller
}

// @router /session2 [get]
func (this *Session2Controller) Get()  {
	// 设置Session
	this.SetSession("userName", "dragonhht")
	this.TplName = "index.html"
}