package models

import (
	"time"
)

type CloudTenantAlarm struct {
	Id        int    `orm:"pk;auto;unique;column(id)" json:"id"`
	UIN       string `orm:"column(uin);unique" json:"uin"`
	Appid     string `orm:"column(appid)" json:"appid"`
	OwnerUin  string `orm:"column(owner_uin)" json:"owner_uin"`
	Account   string `orm:"column(account);unique" json:"account"`
	Title     string `orm:"column(title)" json:"title"`
	LoginURL  string `orm:"column(login_url)" json:"login_url"`
	SecretId  string `orm:"column(secret_id)" json:"secret_id"`
	SecretKey string `orm:"column(secret_key)" json:"secret_key"`

	//
	ProjectId int `orm:"column(project_id)" json:"project_id"`

	// CreateTime 创建时间 .
	CreateTime time.Time `orm:"type(datetime);column(create_time);auto_now_add" json:"create_time"`
	ModifyTime time.Time `orm:"type(datetime);column(modify_time);null;auto_now" json:"modify_time"`
	// 0-enable 1-disable
	Disable int `orm:"disable" json:"disable"`
}
