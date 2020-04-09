package router

import (
	"gin_vue_bs/controller"

	"github.com/gin-gonic/gin"
)

//CollectRoute .
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	return r
}
