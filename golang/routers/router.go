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
			// get all sites
			beego.NSRouter("/", &v1.DownloadController{}, "post:DownloadFiles"),
			beego.NSRouter("/files", &v1.DownloadController{}, "get:AllFiles"),
		),
	)

	// register namespace
	beego.AddNamespace(nsv1)
}
