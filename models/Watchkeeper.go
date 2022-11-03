package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/mindoc-org/mindoc/conf"
)

// WatchKeepering Role
type WKRole uint

const (
	// 运行值班
	WKROperation WKRole = iota
	// 值班经理
	WKRManager
	// 值班团队长
	WKRLeader
)

// Watchkeeping Classification
type WKClassification uint

const (
	// 应用
	WKCApplication WKClassification = iota
	// 平台租户侧
	WKCPlatformTenant
	// 平台网络
	WKCPlatformNetwork
	// 平台操作系统
	WKCPlatformOS
	// 平台数据库
	WKCPlatformDatabase
	// 平台中间件
	WKCPlatformMidware

	// 一线值班、二线值班、值班团队长
	// 
)

type Watchkeeper struct {
	Id             int              `orm:"pk;auto;unique;column(id)" json:"id"`
	Staff          *Member          `orm:"rel(fk)" json:"staff"`
	Role           WKRole           `orm:"type(uint);column(role)" json:"role"`
	Classification WKClassification `orm:"type(uint);column(classification)" json:"classification"`
	// Titile 
	StartTime      time.Time        `orm:"type(datetime);column(start_time);null;index" json:"start_time"`
	EndTime        time.Time        `orm:"type(datetime);column(end_time);null;index" json:"end_time"`
	IsRemote       bool             `orm:"type(bool);column(is_remote)" json:"is_remote"`
	IsLongTerm     bool             `orm:"type(bool);column(is_long_term)" json:"is_long_term"`
	CreateTime     time.Time        `orm:"type(datetime);column(create_time);auto_now_add" json:"create_time"`
	CreateAt       *Member          `orm:"rel(fk)" json:"-"`
	ModifyTime     time.Time        `orm:"column(modify_time);type(datetime);auto_now" json:"modify_time"`
	ModifyAt       *Member          `orm:"rel(fk)" json:"-"`
}

// TableName 获取对应数据库表名.
func (wk *Watchkeeper) TableName() string {
	return "watchkeeper"
}

func (wk *Watchkeeper) TableEngine() string {
	return "INNODB"
}

func (wk *Watchkeeper) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + wk.TableName()
}

func NewWatchkeeper() *Watchkeeper {
	return &Watchkeeper{}
}

func (wk *Watchkeeper) FindById() {
	o := orm.NewOrm()

	user := Watchkeeper{}

	// err := o.Read(&user)
	err := o.QueryTable("md_watchkeeper").Filter("id", 2).RelatedSel().One(&user)

	if err == orm.ErrNoRows {
		log.Println("查询不到")
	} else if err == orm.ErrMissPK {
		log.Println("找不到主键")
	} else {
		log.Println("debug ", user.Id, user.Staff.RealName)
	}
}

func (wk *Watchkeeper) findWKByTime(t time.Time, role WKRole, classify WKClassification) ([]*Watchkeeper, error) {
	o := orm.NewOrm()

	var wk1 []*Watchkeeper
	var wk2 []*Watchkeeper

	WKQuerySetter := o.QueryTable("md_watchkeeper").Filter("role", role).Filter("classification", classify)
	num1, err := WKQuerySetter.Filter("start_time__lte", t).Filter("end_time__gt", t).RelatedSel().All(&wk1)
	if err != nil {
		return nil, err
	}
	num2, err := WKQuerySetter.Filter("is_long_term", true).RelatedSel().All(&wk2)
	if err != nil {
		return nil, err
	}
	if num1 == 0 {
		if num2 == 0 {
			return nil, nil
		} else {
			return wk2, nil
		}
	} else {
		if num2 > 0 {
			wk1 = append(wk1, wk2...)
		}
		return wk1, nil
	}
}

// // time
// func FindWatchkeeperByTime(t time.Time) {

// 	// 一线运行值班
// 	AppPrimary, err := findWKByTime(t, WKROperation, WKCApplication)
// 	// 一线应用经理
// 	AppPrimaryManager, err := findWKByTime(t, WKRManager, WKCApplication)
// 	// 值班团队长
// 	AppPrimaryManager, err := findWKByTime(t, WKRLeader, WKCApplication)

// }
