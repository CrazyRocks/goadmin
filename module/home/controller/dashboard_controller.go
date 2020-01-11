package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goadmin/utils/base"
)

type DashboardController struct {
	base.BaseRouter
}

func (controller *DashboardController) Index(r *ghttp.Request) {
	base.WriteTpl(r, "home/index.html", g.Map{})
}
