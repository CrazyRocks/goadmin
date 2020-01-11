/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  sys_user_controller
* @Version: 1.0.0
* @Date: 2019-11-17
 */

package controller

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"goadmin/module/sys/model"
	"goadmin/utils/base"
	"goadmin/utils/tools"
)

type SysUserController struct {
	base.BaseRouter
}

var (
	controllerNameSysUser = "SysUserController"
)

//SysUser页面信息
func (controller *SysUserController) Index(r *ghttp.Request) {
	base.WriteTpl(r, "sys/user.html", g.Map{})
}

//获取SysUser单条信息
func (controller *SysUserController) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	model := model.SysUser{Id: id}.Get()
	if model.Id <= 0 {
		base.Fail(r, controllerNameSysUser+" get fail")
	}
	base.Succ(r, model)
}

//根据id或者ids删除{.TplModelName}
func (controller *SysUserController) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	for _, id := range ids {
		model := model.SysUser{Id: id}
		model.Delete()
	}
	base.Succ(r, "")
}

//创建一条{.TplModelName}
func (controller *SysUserController) Save(r *ghttp.Request) {
	model := model.SysUser{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(controllerNameSysUser+" save struct error", err)
		base.Error(r, "save error")
	}
	model.Salt = gconv.String(tools.GetId(1))
	if model.Password == "" {
		base.Error(r, "请输入密码")
	}
	model.Password, err = gmd5.Encrypt(model.Password + model.Salt)
	if err != nil {
		glog.Error(controllerNameSysUser+" save struct error", err)
		base.Error(r, "加密错误")
	}
	var num int64
	if model.Id <= 0 {
		num = model.Insert()
	} else {
		num = model.Update()
	}

	if num <= 0 {
		base.Fail(r, controllerNameSysUser+" save fail")
	}

	base.Succ(r, "")
}

//更新一条{.TplModelName}
func (controller *SysUserController) Update(r *ghttp.Request) {
	model := model.SysUser{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(controllerNameSysUser+" save struct error", err)
		base.Error(r, "save error")
	}
	model.Salt = gconv.String(tools.GetId(1))
	if model.Password == "" {
		base.Error(r, "请输入密码")
	}
	model.Password, err = gmd5.Encrypt(model.Password + model.Salt)
	if err != nil {
		glog.Error(controllerNameSysUser+" save struct error", err)
		base.Error(r, "加密错误")
	}
	var num int64
	if model.Id <= 0 {
		num = model.Insert()
	} else {
		num = model.Update()
	}

	if num <= 0 {
		base.Fail(r, controllerNameSysUser+" save fail")
	}

	base.Succ(r, "")
}

//分页列表{.TplModelName}
func (controller *SysUserController) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetQueryMap())
	model := model.SysUser{}
	page := model.Page(&form)
	base.Succ(r, g.Map{"list": page, "form": form})
}
