package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gsession"
	"goadmin/middleware/rtoken"
	"goadmin/module/public/controller"
	"goadmin/utils/base"
	"time"
)

func init() {
	InitConfig()
	InitModules()
	// 启动rtoken
	base.Token = &rtoken.RfToken{
		Timeout:          100 * 1000,
		CacheMode:        2,
		LoginPath:        "/login/submit",
		LoginBeforeFunc:  controller.LoginSubmit,
		LoginAfterFunc:   controller.LoginAfter,
		LogoutPath:       "/user/logout",
		LogoutBeforeFunc: controller.LogoutBefore,
		AuthPaths:        g.SliceStr{"/dashboard", "/sys/*"},
	}
	base.Token.Start()
}

// 用于应用初始化。
func InitConfig() {
	c := g.Config()
	s := g.Server()
	v := g.View()
	//session存内存
	_ = s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time.Minute,
		"SessionStorage": gsession.NewStorageMemory(),
	})

	// 配置对象及视图对象配置
	_ = c.AddPath("config")
	v.SetDelimiters("${", "}")
	_ = v.AddPath("template")

	s.SetServerRoot("public")
	// glog配置
	logpath := c.GetString("setting.logpath")
	_ = glog.SetPath(logpath)
	glog.SetStdoutPrint(true)

	s.SetLogPath(logpath)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(false)
	s.EnableAdmin()
	s.SetPort(8192)
}
