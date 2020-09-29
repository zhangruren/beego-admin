package controllers

import (
	"beego-admin/services"
)

type AdminUserController struct {
	baseController
}

//控制器，初始化函数，基础自父控制器
func (this *AdminUserController) NestPrepare() {
	//fmt.Println("AdminUserController NestPrepare")
}

func (this *AdminUserController)Index()  {

	var adminUserService services.AdminUserService
	this.Data["data"] = adminUserService.GetAllData()

	this.Layout = "public/base.html"
	this.TplName = "admin_user/index.html"
}