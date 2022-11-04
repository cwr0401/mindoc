package main

import (
	"github.com/mindoc-org/mindoc/models"

	"github.com/beego/beego/v2/client/orm"
	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// need to register models in init
	orm.RegisterModel(new(models.CloudTenantAccount))
	orm.RegisterModel(new(models.CloudTenantBasicAlarm))

	// need to register default database
	orm.RegisterDataBase("default", "mysql", "mindoc:MindocPass@tcp(127.0.0.1:3306)/mindoc?charset=utf8")
}

func main() {
	// automatically build table
	// orm.RunSyncdb("default", false, true)

	// create orm object
	_ = orm.NewOrm()

	// insert data
	// o.Insert(user)
}
