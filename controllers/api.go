package controllers

import (
	"github.com/kdada/tinygo"
)

type ApiController struct {
	tinygo.Controller
}

type Content struct {
	Xa string
	K  int
	Io float32
}

func (this *ApiController) Index() {
	tinygo.Debug(this.Context.ParamString("id"))
	this.Api(&Content{
		"xxx",
		1223,
		122.993,
	})
}
