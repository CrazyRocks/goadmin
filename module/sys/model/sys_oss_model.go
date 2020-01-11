/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  sys_oss_model
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
)

import (
	"github.com/gogf/gf/os/gtime"
)

type SysOss struct {
	Id         int         `orm:"id,primary"  json:"Id" gconv:"Id,omitempty"`
	Url        string      `orm:"url"         json:"Url" gconv:"Url,omitempty"`
	CreateTime *gtime.Time `orm:"create_time" json:"CreateTime" gconv:"CreateTime,omitempty"`
	UpdateTime *gtime.Time `orm:"update_time" json:"UpdateTime" gconv:"UpdateTime,omitempty"`
	Status     int         `orm:"status"      json:"Status" gconv:"Status,omitempty"`
}

//创建mSysOss
func (model *SysOss) Insert() int64 {
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
	return res
}

//删除SysOss
func (model SysOss) Delete() int64 {
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
	return res
}

//更新SysOss
func (model *SysOss) Update() int64 {
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
	return res
}

//根据ID查询SysOss
func (model SysOss) Get() SysOss {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return SysOss{}
	}
	var resData SysOss
	err := model.dbModel("t").Where(" id = ? and status=1", model.Id).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysOss{}
	}
	return resData
}

//分页查询SysOss
func (model SysOss) Page(form *base.BaseForm) []SysOss {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
		return []SysOss{}
	}
	where := " status= 1 "
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
		return []SysOss{}
	}
	var resData []SysOss
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	dbModel := model.dbModel("t").Fields(model.columns()).Fields(model.columns())
	err = dbModel.Where(where, params).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" page list error", err)
		return []SysOss{}
	}

	return resData
}

//返回数据库SysOss
func (model SysOss) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB("center").Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

//返回主键SysOss
func (model SysOss) PkVal() int {
	return model.Id
}

//表名SysOss
func (model SysOss) TableName() string {
	return "sys_oss"
}

//列名SysOss
func (model SysOss) columns() string {
	sqlColumns := "t.id as Id,t.url as Url,t.create_time as CreateTime,t.update_time as UpdateTime,t.status as Status"
	return sqlColumns
}
