/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  sys_oss_controller
* @Version: 1.0.0
* @Date: 2019-11-17
 */

package controller

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"novelcenter/module/sys/model"
	"novelcenter/utils/base"
	"novelcenter/utils/tools"
)

type SysOssController struct {
	base.BaseRouter
}

var (
	controllerNameSysOss = "SysOssController"
	ossConfigKey         = "CLOUD_STORAGE_CONFIG_KEY"
)

//SysOss页面信息
func (controller *SysOssController) Index(r *ghttp.Request) {
	base.WriteTpl(r, "sys/oss.html", g.Map{})
}

//获取SysOss单条信息
func (controller *SysOssController) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	model := model.SysOss{Id: id}.Get()
	if model.Id <= 0 {
		base.Fail(r, controllerNameSysOss+" get fail")
	}
	base.Succ(r, model)
}

//根据id或者ids删除{.TplModelName}
func (controller *SysOssController) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	for _, id := range ids {
		model := model.SysOss{Id: id}
		model.Delete()
	}
	base.Succ(r, "")
}

//创建一条{.TplModelName}
func (controller *SysOssController) Save(r *ghttp.Request) {
	model := model.SysOss{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(controllerNameSysOss+" save struct error", err)
		base.Error(r, "save error")
	}
	var num int64
	if model.Id <= 0 {
		num = model.Insert()
	} else {
		num = model.Update()
	}

	if num <= 0 {
		base.Fail(r, controllerNameSysOss+" save fail")
	}

	base.Succ(r, "")
}

//更新一条{.TplModelName}
func (controller *SysOssController) Update(r *ghttp.Request) {
	model := model.SysOss{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(controllerNameSysOss+" save struct error", err)
		base.Error(r, "save error")
	}
	var num int64
	if model.Id <= 0 {
		num = model.Insert()
	} else {
		num = model.Update()
	}

	if num <= 0 {
		base.Fail(r, controllerNameSysOss+" save fail")
	}

	base.Succ(r, "")
}

//分页列表{.TplModelName}
func (controller *SysOssController) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetQueryMap())
	model := model.SysOss{}
	page := model.Page(&form)
	base.Succ(r, g.Map{"list": page, "form": form})
}

//获取OSS的Config
func (controller *SysOssController) GetConfig(r *ghttp.Request) {
	model := model.SysConfig{ParamKey: ossConfigKey}.GetByKey()
	if model.Id <= 0 {
		base.Fail(r, controllerNameSysConfig+" get fail")
	}
	base.Succ(r, model)
}

//保存OSS的Config
func (controller *SysOssController) SaveConfig(r *ghttp.Request) {
	model := model.SysConfig{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(controllerNameSysOss+" save struct error", err)
		base.Error(r, "save error")
	}

	model.ParamKey = ossConfigKey
	model.Remark = "云存储配置信息"
	model.Status = 0
	var num int64
	if model.Id <= 0 {
		num = model.Insert()
	} else {
		num = model.UpdateByKey()
	}
	if num <= 0 {
		base.Fail(r, controllerNameSysConfig+" save fail")
	}

	base.Succ(r, "")
}

//上传文件
func (controller *SysOssController) Upload(r *ghttp.Request) {
	if f, h, e := r.FormFile("upload-file"); e == nil {
		defer f.Close()
		ossConfig := tools.GetCache(ossConfigKey)
		if ossConfig == nil || ossConfig == "" {
			model := model.SysConfig{ParamKey: ossConfigKey}.GetByKey()
			if model.Id <= 0 {
				base.Fail(r, controllerNameSysConfig+" get fail")
			}
			ossConfig = model.ParamValue
		}
		configModel := model.SysOssCloud{}
		err := gjson.DecodeTo(ossConfig, &configModel)
		if err != nil {
			base.Fail(r, "解析错误")
		}
		client, err := oss.New(configModel.AliyunEndPoint, configModel.AliyunAccessKeyId, configModel.AliyunAccessKeySecret)
		if err != nil {
			base.Fail(r, "上传配置错误")
		}
		bucket, err := client.Bucket(configModel.AliyunBucketName)
		if err != nil {
			base.Fail(r, "选择bucket错误")
		}
		name := "file" + gfile.Separator + gconv.String(tools.GetId(2)) + "." + gfile.ExtName(h.Filename)
		err = bucket.PutObject(name, f)
		if err != nil {
			base.Fail(r, "存储错误")
		}
		model := model.SysOss{}
		model.Url = g.Config().GetString("setting.cdn") + name
		var num int64
		num = model.Insert()
		if num <= 0 {
			base.Fail(r, "存入数据库失败")
		}
		base.Succ(r, model)
	} else {
		r.Response.Write(e.Error())
	}
}
