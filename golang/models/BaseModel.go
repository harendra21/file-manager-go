package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	runmode, _ := config.String("runmode")
	dbuser, _ := config.String("dbuser")
	dbpass, _ := config.String("dbpass")
	dbname, _ := config.String("dbname")
	dbhost, _ := config.String("dbhost")
	dbport, _ := config.String("dbport")

	dbUser := dbuser
	dbPwd := dbpass
	dbName := dbname
	dbHost := dbhost
	dbPort := dbport
	dbString := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&sslmode=disable"
	// Register Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// Register default database
	orm.RegisterDataBase("default", "mysql", dbString)
	orm.RegisterModel(new(User), new(Site), new(Detail))
	if runmode == "dev" {
		orm.Debug = true
	}
}
