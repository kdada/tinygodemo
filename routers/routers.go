package routers

import (
	"tinygodemo/controllers"

	"github.com/kdada/tinygo"
)

//注册路由
func init() {
	tinygo.RootRouter.AddChild(tinygo.NewControllerRouter(new(controllers.HomeController)))
}
