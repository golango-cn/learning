package router

import (
	"github.com/gin-gonic/gin"
	"ndb/server/controller"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.POST("/login", controller.Login)
	r.GET("/menus", controller.GetMenus)
	r.GET("/users", controller.GetUsers)
	r.POST("/rights", controller.GetRights)
	r.POST("/roles", controller.GetRoles)
	r.GET("/rules", controller.GetRules)

	r.GET("/good/cates", controller.GetCates)
	r.POST("/good/attrs/:attr_type", controller.GetAttrs)

	return r

}
