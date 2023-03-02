// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego/controllers"
	v1 "beego/controllers/v1"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	// init namespace
	nsv1 := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/download",
			beego.NSRouter("/", &v1.DownloadController{}, "post:DownloadFiles"),
		),
		beego.NSNamespace("/file",
			beego.NSRouter("/", &v1.FileController{}, "get:AllFiles"),
			beego.NSRouter("/create", &v1.FileController{}, "get:Create"),
			beego.NSRouter("/move", &v1.FileController{}, "get:MoveRename"),
			beego.NSRouter("/copy", &v1.FileController{}, "get:Copy"),
			beego.NSRouter("/delete", &v1.FileController{}, "get:Delete"),
			beego.NSRouter("/rename", &v1.FileController{}, "get:MoveRename"),
			beego.NSRouter("/zip", &v1.FileController{}, "get:Zip"),
			beego.NSRouter("/unzip", &v1.FileController{}, "get:Unzip"),
		),
	)

	// register namespace
	beego.AddNamespace(nsv1)
}
