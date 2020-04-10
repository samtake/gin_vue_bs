package common

import (
	"fmt"

	"gin_vue_bs/model"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

//DB .
var DB *gorm.DB

//InitDB .
func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	passwoed := viper.GetString("datasource.password")
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
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

//GetDB .
func GetDB() *gorm.DB {
	return DB
}
