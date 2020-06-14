package service

import (
	"log"
	"webDemo1/config"
	"webDemo1/model"
)

//添加
func InsertUser(uname string,upassword string)bool  {
	var user=model.User{
		Uname:uname,
		Upassword: upassword,
	}
	err:=config.DB.Create(&user).Error
	if err!=nil{
		log.Fatal("添加错误",err)
		return false
	}
	return true
}
//查询
func CheckUser(uname string,upassword string)int {
	var user=model.User{
		Uname:uname,
		Upassword: upassword,
	}
	var num int
	err:=config.DB.Find(&user).Count(&num).Error
	if err!=nil{
		log.Fatal("查询错误",err)
		return -1
	}
	return num
}
