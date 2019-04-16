package routers

import (
	"beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    /*beego.Router("/user/get", &controllers.UserController{})*/
    beego.Include(&controllers.UserController{})
    beego.Include(&controllers.SessionController{})
    beego.Include(&controllers.Session2Controller{})
}
