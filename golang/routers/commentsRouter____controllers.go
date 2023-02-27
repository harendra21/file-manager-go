package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego/controllers:SiteController"] = append(beego.GlobalControllerRouter["beego/controllers:SiteController"],
        beego.ControllerComments{
            Method: "AddNew",
            Router: "/",
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:SiteController"] = append(beego.GlobalControllerRouter["beego/controllers:SiteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/api/v1/site",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
