/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  sys_config_model
* @Version: 1.0.0
* @Date: 2019-11-17
 */
package model

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"goadmin/utils/base"
	"goadmin/utils/tools"
)

import (
	"github.com/gogf/gf/os/gtime"
)

type SysConfig struct {
	Id         int         `orm:"id,primary"       json:"Id" gconv:"Id,omitempty"`
	ParamKey   string      `orm:"param_key,unique" json:"ParamKey" gconv:"ParamKey,omitempty"`
	ParamValue string      `orm:"param_value"      json:"ParamValue" gconv:"ParamValue,omitempty"`
	Remark     string      `orm:"remark"           json:"Remark" gconv:"Remark,omitempty"`
	CreateTime *gtime.Time `orm:"create_time"      json:"CreateTime" gconv:"CreateTime,omitempty"`
	UpdateTime *gtime.Time `orm:"update_time"      json:"UpdateTime" gconv:"UpdateTime,omitempty"`
	Status     int         `orm:"status"           json:"Status" gconv:"Status,omitempty"`
}

//创建mSysConfig
func (model *SysConfig) Insert() int64 {
	model.CreateTime = gtime.Now()
	model.UpdateTime = gtime.Now()
	model.Status = 1
	r, err := model.dbModel().Data(model).Insert()
	if err != nil {
		glog.Error(model.TableName()+" insert error", err)
		return 0
	}

	res, err2 := r.RowsAffected()
	if err2 != nil {
		glog.Error(model.TableName()+" insert res error", err2)
		return 0
	} else if res > 0 {
		lastId, err2 := r.LastInsertId()
		if err2 != nil {
			glog.Error(model.TableName()+" LastInsertId res error", err2)
			return 0
		} else {
			model.Id = gconv.Int(lastId)
		}
	}
	tools.SetCache(model.ParamKey, model.ParamValue)
	return res
}

//删除SysConfig
func (model SysConfig) Delete() int64 {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " delete id error")
		return 0
	}
	r, err := model.dbModel().Where(" id = ?", model.Id).Delete()
	if err != nil {
		glog.Error(model.TableName()+" delete error", err)
		return 0
	}
	res, err2 := r.RowsAffected()
	if err2 != nil {
		glog.Error(model.TableName()+" delete res error", err2)
		return 0
	}
	tools.DelCache(model.ParamKey)
	return res
}

//更新SysConfig
func (model *SysConfig) Update() int64 {
	model.UpdateTime = gtime.Now()
	r, err := model.dbModel().Data(model).Where(" id = ?", model.Id).Update()
	if err != nil {
		glog.Error(model.TableName()+" update error", err)
		return 0
	}
	res, err2 := r.RowsAffected()
	if err2 != nil {
		glog.Error(model.TableName()+" update res error", err2)
		return 0
	}
	tools.SetCache(model.ParamKey, model.ParamValue)
	return res
}

func (model *SysConfig) UpdateByKey() int64 {
	model.UpdateTime = gtime.Now()
	r, err := model.dbModel().Data(model).Where(" param_key = ?", model.ParamKey).Update()
	if err != nil {
		glog.Error(model.TableName()+" update error", err)
		return 0
	}
	res, err2 := r.RowsAffected()
	if err2 != nil {
		glog.Error(model.TableName()+" update res error", err2)
		return 0
	}
	tools.SetCache(model.ParamKey, model.ParamValue)
	return res
}

//根据Key查询SysConfig
func (model SysConfig) GetByKey() SysConfig {
	if model.ParamKey == "" {
		glog.Error("根据KEY获取配置失败")
		return SysConfig{}
	}
	var resData SysConfig
	err := model.dbModel("t").Where(" param_key = ?", model.ParamKey).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysConfig{}
	}

	return resData
}

//根据ID查询SysConfig
func (model SysConfig) Get() SysConfig {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return SysConfig{}
	}
	var resData SysConfig
	err := model.dbModel("t").Where(" id = ? and status=1", model.Id).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysConfig{}
	}
	return resData
}

//分页查询SysConfig
func (model SysConfig) Page(form *base.BaseForm) []SysConfig {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
		return []SysConfig{}
	}
	where := " status = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	num, err := model.dbModel("t").Where(where, params).Count()
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []SysConfig{}
	}
	var resData []SysConfig
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	dbModel := model.dbModel("t").Fields(model.columns()).Fields(model.columns())
	err = dbModel.Where(where, params).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" page list error", err)
		return []SysConfig{}
	}

	return resData
}

//返回数据库SysConfig
func (model SysConfig) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB("center").Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

//返回主键SysConfig
func (model SysConfig) PkVal() int {
	return model.Id
}

//表名SysConfig
func (model SysConfig) TableName() string {
	return "sys_config"
}

//列名SysConfig
func (model SysConfig) columns() string {
	sqlColumns := "t.create_time as CreateTime,t.update_time as UpdateTime,t.status as Status,t.id as Id,t.param_key as ParamKey,t.param_value as ParamValue,t.remark as Remark"
	return sqlColumns
}
