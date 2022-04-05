package mysql

import (
	"fmt"

	"github.com/hollson/kendo/config"
	"github.com/hollson/kendo/infrastructure/repo"
	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"
)

var readEngine *xorm.Engine
var writeEngine *xorm.Engine

// InitDB 初始化DB引擎
func InitDB() {
	var err error
	readEngine, err = xorm.NewEngine("mysql", config.MysqlSource)
	if err != nil {
		panic(fmt.Errorf("mysql error: %v", err))
	}

	readEngine.SetMaxIdleConns(config.MysqlMaxIdle)
	readEngine.SetMaxOpenConns(config.MysqlMaxOpen)

	// 读写需要分开不同的库
	writeEngine = readEngine
	if err != nil {
		panic(fmt.Errorf("mysql error: %v", err))
	}
	writeEngine.SetMaxIdleConns(config.MysqlMaxIdle)
	writeEngine.SetMaxOpenConns(config.MysqlMaxOpen)

	if config.MysqlShowSQL {
		readEngine.ShowSQL(true)
		writeEngine.ShowSQL(true)
	}

	writeEngine.Sync2(new(repo.FileInfo))
}
