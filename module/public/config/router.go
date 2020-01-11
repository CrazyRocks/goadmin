package config

import (
	"github.com/gogf/gf/frame/g"
	"goadmin/module/public/controller"
)

func InitRouter() {
	urlPath := g.Config().GetString("url-path")
	s := g.Server()
	// 首页
	s.BindHandler(urlPath+"/", controller.Login)
}
