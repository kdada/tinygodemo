package routers

import (
	"fmt"
	"tinygodemo/controllers"

	"github.com/kdada/tinygo"
	"github.com/kdada/tinygo/router"
)

//注册路由
func init() {
	tinygo.AddRouter(tinygo.NewControllerRouter(new(controllers.HomeController)))
	tinygo.AddRouter(tinygo.NewComplexControllerRouter("/xx/ssdd",
		new(controllers.ApiController),
		tinygo.RouterInfos{
			tinygo.NewGetRouterInfo("Index", nil),
		}))
	tinygo.AddRouter(tinygo.NewComplexControllerRouter("/xx/errr",
		new(controllers.ApiController),
		tinygo.RouterInfos{
			tinygo.NewGetRouterInfo("Index", tinygo.RouterParams{
				`list_(id=\d+).html`,
			}),
		}))
	tinygo.AddRouter(tinygo.NewRestfulControllerRouter(
		new(controllers.RestfulTestController),
	))
	tinygo.AddRouter(tinygo.NewPathRestfulControllerRouter("/some/sss", new(controllers.RestfulTestController)))
	tinygo.AddRouter(tinygo.NewFunctionRouter(`/xe/tend(kk=\d+)`, func(context router.RouterContext) {
		var ctx = context.(*tinygo.HttpContext)
		fmt.Println(ctx.ParamString("kk"))
		tinygo.HttpNotFound(ctx.ResponseWriter(), ctx.Request())
	}))
}
