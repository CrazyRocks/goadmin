/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  sys_user_model
* @Version: 1.0.0
* @Date: 2019-11-04
 */
package model

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

import (
	"github.com/gogf/gf/os/gtime"
)

type SysLoginUser struct {
	Id         int         `orm:"id,primary"      json:"Id" gconv:"Id,omitempty"`
	Username   string      `orm:"username,unique" json:"Username" gconv:"Username,omitempty"`
	Password   string      `orm:"password"        json:"Password" gconv:"Password,omitempty"`
	Salt       string      `orm:"salt"            json:"Salt" gconv:"Salt,omitempty"`
	Mobile     string      `orm:"mobile"          json:"Mobile" gconv:"Mobile,omitempty"`
	CreateTime *gtime.Time `orm:"create_time"     json:"CreateTime" gconv:"CreateTime,omitempty"`
	UpdateTime *gtime.Time `orm:"update_time"     json:"UpdateTime" gconv:"UpdateTime,omitempty"`
	Status     int         `orm:"status"          json:"Status" gconv:"Status,omitempty"`
}

//根据手机号查询系统用户
func (model SysLoginUser) GetUserByMobile() SysLoginUser {
	var resData SysLoginUser
	err := model.dbModel("t").Where("mobile = ?", model.Mobile).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysLoginUser{}
	}
	return resData

}

//返回数据库SysUser
func (model SysLoginUser) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB("center").Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

//返回主键SysUser
func (model SysLoginUser) PkVal() int {
	return model.Id
}

//表名SysUser
func (model SysLoginUser) TableName() string {
	return "sys_user"
}

//列名SysUser
func (model SysLoginUser) columns() string {
	sqlColumns := "t.id as Id,t.username as Username,t.password as Password,t.salt as Salt,t.mobile as Mobile,t.create_time as CreateTime,t.update_time as UpdateTime,t.status as Status"
	return sqlColumns
}
