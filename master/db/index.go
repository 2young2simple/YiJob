package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/2young2simple/YiJob/model"
)

func init() {
	sql := beego.AppConfig.String("mysql")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", sql)
	orm.RegisterModel(new(model.JobModel))
	orm.RunSyncdb("default", false, true)
}
