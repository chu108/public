package mysqldb

import (
	"fmt"

	"github.com/xie1xiao1jun/public/dev"

	"github.com/xie1xiao1jun/public/mylog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MySqlDB struct {
	*gorm.DB
	IsInit bool
}

func OnInitDBOrm(dataSourceName string) (orm *MySqlDB) {
	orm = new(MySqlDB)
	orm.OnGetDBOrm(dataSourceName)
	return
}

func (i *MySqlDB) OnGetDBOrm(dataSourceName string) (orm *gorm.DB) {
	if i.DB == nil {
		var err error
		i.DB, err = gorm.Open("mysql", dataSourceName)
		if err != nil {
			mylog.Print(mylog.Log_Error, fmt.Sprintf("Got error when connect database, '%v'", err))
			return nil
		}
		i.IsInit = true
	}

	i.DB.SingularTable(true) //全局禁用表名复数
	if dev.OnIsDev() {
		i.DB.LogMode(true)
		//beedb.OnDebug = true
	} else {
		i.DB.SetLogger(DbLog{})
	}
	orm = i.DB
	return
}

func (i *MySqlDB) OnDestoryDB() {
	if i.DB != nil {
		i.DB.Close()
		i.DB = nil
	}
}
