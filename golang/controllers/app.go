package controllers

import (
	"beego/lib"

	"errors"

	beego "github.com/beego/beego/v2/server/web"
)

type AppController struct {
	beego.Controller
}

type Error struct {
	Status int    `json:"status"`
	Code   int    `json:"error_code"`
	Error  string `json:"error_msg"`
}

type Success struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (a *AppController) Response(res interface{}) {
	a.Data["json"] = Success{Status: 1, Data: res}
	err := a.ServeJSON()
	if err != nil {
		lib.Logger{}.Error(err)
	}
}

func (a *AppController) ThrowError(code int, msg ...string) {

	errMsg := ""
	if len(msg) > 0 {
		errMsg = msg[0]
	} else {
		errMsg = errs[code]
	}

	a.Data["json"] = Error{Status: 0, Code: code, Error: errMsg}
	err := a.ServeJSON()
	if err != nil {
		lib.Logger{}.Error(err)
	}

	if code == 1001 {
		metaData := map[string]interface{}{
			"url":    a.Ctx.Request.RequestURI,
			"method": a.Ctx.Request.Method,
		}
		lib.Logger{}.Info(errMsg, metaData)
	} else {
		lib.Logger{}.Error(errors.New(errMsg))
	}
}

var errs map[int]string = map[int]string{
	1001: "Resource not found",
	1002: "",
	1003: "Invalid Params",
	1004: "Resource creation error",
	1005: "Can't delete resource",
}
