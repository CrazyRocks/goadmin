/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  sys_user_model
* @Version: 1.0.0
* @Date: 2019-11-17
 */
package model

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"goadmin/utils/base"
)

import (
	"github.com/gogf/gf/os/gtime"
)

type SysUser struct {
	Id         int         `orm:"id,primary"      json:"Id" gconv:"Id,omitempty"`
	Username   string      `orm:"username,unique" json:"Username" gconv:"Username,omitempty"`
	Password   string      `orm:"password"        json:"Password" gconv:"Password,omitempty"`
	Salt       string      `orm:"salt"            json:"Salt" gconv:"Salt,omitempty"`
	Mobile     string      `orm:"mobile"          json:"Mobile" gconv:"Mobile,omitempty"`
	CreateTime *gtime.Time `orm:"create_time"     json:"CreateTime" gconv:"CreateTime,omitempty"`
	UpdateTime *gtime.Time `orm:"update_time"     json:"UpdateTime" gconv:"UpdateTime,omitempty"`
	Status     int         `orm:"status"          json:"Status" gconv:"Status,omitempty"`
}

//创建mSysUser
func (model *SysUser) Insert() int64 {
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

//删除SysUser
func (model SysUser) Delete() int64 {
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

//更新SysUser
func (model *SysUser) Update() int64 {
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

//根据ID查询SysUser
func (model SysUser) Get() SysUser {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return SysUser{}
	}
	var resData SysUser
	err := model.dbModel("t").Where(" id = ? and status=1", model.Id).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysUser{}
	}
	return resData
}

//分页查询SysUser
func (model SysUser) Page(form *base.BaseForm) []SysUser {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
		return []SysUser{}
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
		return []SysUser{}
	}
	var resData []SysUser
	fmt.Println("当前排序:%s", form.OrderBy)
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	dbModel := model.dbModel("t").Fields(model.columns()).Fields(model.columns())
	err = dbModel.Where(where, params).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" page list error", err)
		return []SysUser{}
	}

	return resData
}

//返回数据库SysUser
func (model SysUser) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB("center").Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

//返回主键SysUser
func (model SysUser) PkVal() int {
	return model.Id
}

//表名SysUser
func (model SysUser) TableName() string {
	return "sys_user"
}

//列名SysUser
func (model SysUser) columns() string {
	sqlColumns := "t.id as Id,t.username as Username,t.password as Password,t.salt as Salt,t.mobile as Mobile,t.create_time as CreateTime,t.update_time as UpdateTime,t.status as Status"
	return sqlColumns
}
