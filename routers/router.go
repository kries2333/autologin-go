package routers

import (
	"autologin-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/fileopt", &controllers.FileOptUploadController{})
}
