package controller

import (
	"log"
	"net/http"

	"gin_vue_bs/common"
	"gin_vue_bs/dto"
	"gin_vue_bs/model"
	"gin_vue_bs/response"
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
		response.Response(ctx, http.StatusServiceUnavailable, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	//名称如果没有传，则返回随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	//创建用户
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	DB.Create(&newUser)

	//返回结果
	response.Response(ctx, http.StatusUnprocessableEntity, 200, nil, "注册成功")
}

//Login 登录.
func Login(ctx *gin.Context) {

	DB := common.GetDB()

	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	//手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	//判断密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)

	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统错误")
		log.Printf("token generate error :%s", err)
		return
	}

	response.Response(ctx, http.StatusUnprocessableEntity, 200, gin.H{"token": token}, "登录成功")

}

//Info 用户信息.
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	response.Response(ctx, http.StatusUnprocessableEntity, 200, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
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
