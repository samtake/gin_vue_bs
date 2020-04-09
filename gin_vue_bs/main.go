package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/e421083458/gorm" //https://gorm.io
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(110;not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {

	db := InitDb()
	defer db.Close()

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.POST("/api/auth/register", func(ctx *gin.Context) {
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
			name = RandomString(10)
		}
		//判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity,
				gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}

		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		//返回结果
		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})

	panic(r.Run(":8099"))
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

//返回随机字符串
func RandomString(n int) string {
	var letters = []byte("iasdhjklfhascvxnjklasdfhjkasdfklasdfhnjklasdfjklasdfjklfasdsdfjk")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func InitDb() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "gin_vue_bs"
	username := "root"
	passwoed := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		passwoed,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err" + err.Error())
	}

	//创建数据表
	db.AutoMigrate(&User{})

	return db
}
