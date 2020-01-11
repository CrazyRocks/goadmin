package controller

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"goadmin/middleware/rtoken"
	"goadmin/module/public/model"
	"goadmin/utils/base"
	"goadmin/utils/bean"
)

// Login 登录页面
func Login(r *ghttp.Request) {
	err := r.Response.WriteTpl("public/login.html", g.Map{})

	if err != nil {
		glog.Error(err)
	}
}

// LoginSubmit 登录认证
func LoginSubmit(r *ghttp.Request) (string, interface{}) {
	mobile := r.GetString("mobile")
	password := r.GetString("password")
	if mobile == "" || password == "" {
		base.Fail(r, "手机号或者密码错误")
	}
	sysLoginUser := model.SysLoginUser{Mobile: mobile}.GetUserByMobile()
	if sysLoginUser.Id <= 0 {
		base.Fail(r, "手机号或者密码错误："+mobile)
	}
	if sysLoginUser.Status <= 0 {
		base.Fail(r, "账号状态异常，请联系管理员")
	}

	reqPassword, err := gmd5.Encrypt(password + sysLoginUser.Salt)
	if err != nil {
		glog.Error("login password encrypt error", err)
		base.Error(r, "login password encrypt error")
	}

	if reqPassword != sysLoginUser.Password {
		base.Fail(r, "用户名或者密码错误："+mobile)
	}
	sessionUser := bean.SessionUser{
		Id:       sysLoginUser.Id,
		Username: sysLoginUser.Username,
		Mobile:   sysLoginUser.Mobile,
	}
	return mobile, sessionUser
}

func LoginAfter(r *ghttp.Request, respData rtoken.Resp) {
	r.Cookie.Set("token", respData.GetString("token"))
	if !respData.Success() {
		r.Response.WriteJson(respData)
	} else {
		r.Response.WriteJson(rtoken.Succ(g.Map{
			"token": respData.GetString("token"),
		}))
	}
}

// 登出
func LogoutBefore(r *ghttp.Request) bool {
	mobile := base.GetUser(r).Mobile
	sysLoginUser := model.SysLoginUser{Mobile: mobile}.GetUserByMobile()
	if sysLoginUser.Mobile != mobile {
		// 登出用户不存在
		glog.Warning("logout mobile error", mobile)
		return false
	}
	r.Cookie.Remove("token")
	return true
}
