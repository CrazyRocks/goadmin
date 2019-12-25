/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  router
* @Version: 1.0.0
* @Date: 2019-11-17
 */
package config

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"novelcenter/module/sys/controller"
)

func InitRouter() {
	urlPath := g.Config().GetString("url-path")
	s := g.Server()
	s.Group(urlPath+"/sys", func(g *ghttp.RouterGroup) {
		sysConfigController := new(controller.SysConfigController)
		g.ALL("/config", sysConfigController)
		g.POST("/config/page", sysConfigController.Page)
		g.GET("/config/get/{id}", sysConfigController.Get)
		g.POST("/config/save", sysConfigController.Save)
		g.POST("/config/update", sysConfigController.Update)
		g.POST("/config/delete", sysConfigController.Delete)

		sysOssController := new(controller.SysOssController)
		g.ALL("/oss", sysOssController)
		g.POST("/oss/page", sysOssController.Page)
		g.GET("/oss/get/{id}", sysOssController.Get)
		g.POST("/oss/save", sysOssController.Save)
		g.POST("/oss/update", sysOssController.Update)
		g.POST("/oss/delete", sysOssController.Delete)
		g.GET("/oss/config/get", sysOssController.GetConfig)
		g.POST("/oss/config/save", sysOssController.SaveConfig)
		g.GET("/oss/upload", sysOssController.Upload)

		sysUserController := new(controller.SysUserController)
		g.ALL("/user", sysUserController)
		g.POST("/user/page", sysUserController.Page)
		g.GET("/user/get/{id}", sysUserController.Get)
		g.POST("/user/save", sysUserController.Save)
		g.POST("/user/update", sysUserController.Update)
		g.POST("/user/delete", sysUserController.Delete)

	})
}
