package main

import (
	"github.com/gin-gonic/gin"

	"gin_vue_bs/common"
	"gin_vue_bs/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := common.InitDB()
	defer db.Close()

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r = router.CollectRoute(r)

	panic(r.Run(":8099"))
}
