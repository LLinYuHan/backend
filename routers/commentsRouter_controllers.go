package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Fruit-backend/controllers:ProfileController"] = append(beego.GlobalControllerRouter["Fruit-backend/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Fruit-backend/controllers:ProfileController"] = append(beego.GlobalControllerRouter["Fruit-backend/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Fruit-backend/controllers:ProfileController"] = append(beego.GlobalControllerRouter["Fruit-backend/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Fruit-backend/controllers:ProfileController"] = append(beego.GlobalControllerRouter["Fruit-backend/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Fruit-backend/controllers:ProfileController"] = append(beego.GlobalControllerRouter["Fruit-backend/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["Fruit-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/register/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
