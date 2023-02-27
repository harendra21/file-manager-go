package main

import (
	_ "beego/lib"
	_ "beego/routers"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// dbuser, _ := config.String("dbuser")
	// dbpass, _ := config.String("dbpass")
	// dbname, _ := config.String("dbname")
	// dbhost, _ := config.String("dbhost")
	// dbport, _ := config.String("dbport")

	// dbUser := dbuser
	// dbPwd := dbpass
	// dbName := dbname
	// dbHost := dbhost
	// dbPort := dbport
	// dbString := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	// err := orm.RegisterDriver("mysql", orm.DRMySQL)
	// if err != nil {
	// 	lib.Logger{}.Error(err)
	// }
	// orm.RegisterDataBase("default", "mysql", dbString)

	// // autosync
	// // db alias
	// name := "default"
	// // drop table and re-create
	// force := false

	// // print log
	// verbose := true

	// // error
	// err := orm.RunSyncdb(name, force, verbose)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	beego.Run()
}

// bee generate scaffold myuser -fields="first_name:string:255,last_name:string:255,email:string:255" -driver=mysql -conn="mysql://root:harendra21@localhost:3306/sitepro?sslmode=disable"

// bee generate docs
