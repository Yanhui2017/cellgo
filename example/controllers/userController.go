package controllers

import (
	"library/service"

	"github.com/mrkt/cellgo"
)

type UserController struct {
	cellgo.Controller
}

func (this *UserController) Before() {
	//init service
	this.GetService(&service.UserService{})
}

func (this *UserController) Run() {
	//param1 funcName, param2 funcParam ...
	userInfo := this.GetServiceFunc("GetUserInfo", "tommy").(map[string]string)
	this.Data["Username"] = userInfo["Username"]
	this.Data["Email"] = userInfo["Email"]
	this.Data["URI"] = this.Net.Input.Site() + this.Net.Input.URI()
	this.TplName = "index.html"
}

func (this *UserController) Add() {
	username := this.Net.Input.GetGP("username", true)
	email := this.Net.Input.GetGP("email", true)
	this.Data["Username"] = username
	this.Data["Email"] = email
	this.Data["URI"] = this.Net.Input.Site() + this.Net.Input.URI()
	this.TplName = "index.html"
}

func (this *UserController) Session() {
	if res := this.Net.Input.Session.Get("user"); res != nil {
		user := res.(map[string]string)
		this.Data["Username"] = user["username"]
		this.Data["Email"] = user["email"]
	} else {
		var user map[string]string = make(map[string]string)
		user["username"] = "tommy"
		user["email"] = "tommy.jin@aliyun.com"
		this.Net.Input.Session.Set("user", user)
	}
	this.Data["URI"] = this.Net.Input.Site() + this.Net.Input.URI()
	this.TplName = "index.html"
}

func (this *UserController) Cookie() {
	if res := this.Net.Input.Cookie.Get("user"); res != nil {
		user := res.(map[string]string)
		this.Data["Username"] = user["username"]
		this.Data["Email"] = user["email"]
	} else {
		var user map[string]string = make(map[string]string)
		user["username"] = "tommy"
		user["email"] = "tommy.jin@aliyun.com"
		this.Net.Input.Cookie.Set("user", user)

	}
	this.Data["URI"] = this.Net.Input.Site() + this.Net.Input.URI()
	this.TplName = "index.html"
}
