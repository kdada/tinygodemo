package controllers

import (
	"fmt"

	"github.com/kdada/tinygo"
)

type HomeController struct {
	tinygo.Controller
}

// Routers 返回当前控制器可以使用的方法路由信息
func (this *HomeController) Routers() []interface{} {
	return []interface{}{tinygo.NewRouterInfo("Index", tinygo.HttpMethodGet, nil)}
}

func (this *HomeController) Index() {
	var x, ok = this.Context.Session().Int("test")
	fmt.Println("x:", x, ok, this.Context.Session().SessionId())
	x++
	this.Context.Session().SetInt("test", x)
	this.SimpleView()
}
