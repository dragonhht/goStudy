package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego/controllers:Session2Controller"] = append(beego.GlobalControllerRouter["beego/controllers:Session2Controller"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/session2`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:SessionController"] = append(beego.GlobalControllerRouter["beego/controllers:SessionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/session`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/user/get/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
