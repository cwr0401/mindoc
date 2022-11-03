package models

import (
	"time"

	"github.com/mindoc-org/mindoc/conf"
)

// cooperator

type CloudTenantAccount struct {
	Id        int    `orm:"pk;auto;unique;column(id)" json:"id"`
	UIN       string `orm:"column(uin);unique" json:"uin"`
	Appid     string `orm:"column(appid)" json:"appid"`
	OwnerUin  string `orm:"column(owner_uin)" json:"owner_uin"`
	Account   string `orm:"column(account);unique" json:"account"`
	Title     string `orm:"column(title)" json:"title"`
	LoginURL  string `orm:"column(login_url)" json:"login_url"`
	SecretId  string `orm:"column(secret_id)" json:"secret_id"`
	SecretKey string `orm:"column(secret_key)" json:"secret_key"`

	// CreateTime 创建时间 .
	CreateTime time.Time `orm:"type(datetime);column(create_time);auto_now_add" json:"create_time"`
	ModifyTime time.Time `orm:"type(datetime);column(modify_time);null;auto_now" json:"modify_time"`
	// 0-enable 1-disable
	Disable int `orm:"disable" json:"disable"`
}

// type CloudTenantAccountAkSk struct {

// }

func (account *CloudTenantAccount) IsEnable() bool {
	return account.Disable == 0
}

// TableName 获取对应数据库表名.
func (account *CloudTenantAccount) TableName() string {
	return "cloud_tenant_account"
}

func (account *CloudTenantAccount) TableEngine() string {
	return "INNODB"
}

func (account *CloudTenantAccount) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + account.TableName()
}

func (account *CloudTenantAccount) GetLoginURL() string {
	if account.LoginURL != "" {
		return account.LoginURL
	}
	return "http://ow.yun.ccb.com/login/subAccount/" + account.OwnerUin
}

func NewCloudTenantAccount() *CloudTenantAccount {
	return &CloudTenantAccount{}
}
