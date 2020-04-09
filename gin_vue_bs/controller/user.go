package controller

import (
	"net/http"

	"gin_vue_bs/common"
	"gin_vue_bs/model"
	"gin_vue_bs/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Register 注册.
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": 500, "msg": "系统错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	DB.Create(&newUser)

	//返回结果
	ctx.JSON(200, gin.H{
		"msg":  "注册成功",
		"code": "200",
	})
}

//Login 登录.
func Login(ctx *gin.Context) {

	DB := common.GetDB()

	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	//手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	//判断密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	//发放token
	token := "66"

	ctx.JSON(200, gin.H{
		"msg":  "登录成功",
		"code": "200",
		"data": gin.H{"token": token},
	})

}

//isTelephoneExist .
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
