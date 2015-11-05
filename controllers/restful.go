package controllers

import (
	"github.com/kdada/tinygo"
)

type RestfulTestController struct {
	tinygo.RestfulController
}

// Get HTTP GET对应方法
func (this *RestfulTestController) Get() {
	this.Api(&Content{
		"xnbbxx",
		122223,
		122.99007,
	})
}
