/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  sys_config_controller
* @Version: 1.0.0
* @Date: 2019-11-17
 */

package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"goadmin/module/sys/model"
	"goadmin/utils/base"
)

type SysConfigController struct {
	base.BaseRouter
}

var (
	controllerNameSysConfig = "SysConfigController"
)

//SysConfig页面信息
func (controller *SysConfigController) Index(r *ghttp.Request) {
	base.WriteTpl(r, "sys/config.html", g.Map{})
}

//获取SysConfig单条信息
func (controller *SysConfigController) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	model := model.SysConfig{Id: id}.Get()
	if model.Id <= 0 {
		base.Fail(r, controllerNameSysConfig+" get fail")
	}
	base.Succ(r, model)
}

//根据id或者ids删除{.TplModelName}
func (controller *SysConfigController) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	for _, id := range ids {
		model := model.SysConfig{Id: id}
		model.Delete()
	}
	base.Succ(r, "")
}

//创建一条{.TplModelName}
func (controller *SysConfigController) Save(r *ghttp.Request) {
	model := model.SysConfig{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(controllerNameSysConfig+" save struct error", err)
		base.Error(r, "save error")
	}
	var num int64
	if model.Id <= 0 {
		num = model.Insert()
	} else {
		num = model.Update()
	}

	if num <= 0 {
		base.Fail(r, controllerNameSysConfig+" save fail")
	}

	base.Succ(r, "")
}

//更新一条{.TplModelName}
func (controller *SysConfigController) Update(r *ghttp.Request) {
	model := model.SysConfig{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(controllerNameSysConfig+" save struct error", err)
		base.Error(r, "save error")
	}
	var num int64
	if model.Id <= 0 {
		num = model.Insert()
	} else {
		num = model.Update()
	}

	if num <= 0 {
		base.Fail(r, controllerNameSysConfig+" save fail")
	}

	base.Succ(r, "")
}

//分页列表{.TplModelName}
func (controller *SysConfigController) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetQueryMap())
	model := model.SysConfig{}
	page := model.Page(&form)
	base.Succ(r, g.Map{"list": page, "form": form})
}
