package controller

import (
	"net/http"

	"gin_vue_bs/common"
	"gin_vue_bs/model"
	"gin_vue_bs/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(ctx *gin.Context) {

	DB := common.GetDB()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//验证数据
	if len(telephone) != 11 {
		ctx.JSON(http.StatusServiceUnavailable,
			gin.H{"code": 422, "msg": "手机号必须为11为"})
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": 422, "msg": "密码不能少于6位"})
	}

	//名称如果没有传，则返回随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}

	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	//返回结果
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
