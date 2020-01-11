package config

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goadmin/module/home/controller"
)

func InitRouter() {
	urlPath := g.Config().GetString("url-path")
	s := g.Server()

	s.Group(urlPath+"/dashboard", func(g *ghttp.RouterGroup) {
		// 首页展示
		dashboardController := new(controller.DashboardController)
		g.ALL("dashboard", dashboardController)
		g.GET("/", dashboardController.Index)
	})
}
