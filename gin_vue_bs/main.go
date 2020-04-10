package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"gin_vue_bs/common"
	"gin_vue_bs/router"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r = router.CollectRoute(r)

	panic(r.Run(":8099"))
}

//InitConfig 初始化配置文件.
func InitConfig() {
	workDir, _ := os.Getwd()                 //工程项目目录
	viper.SetConfigName("application")       //配置文件名
	viper.SetConfigType("yml")               //配置文件类型
	viper.AddConfigPath(workDir + "/config") //配置文件目录
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
