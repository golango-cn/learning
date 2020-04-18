package controller

import "github.com/gin-gonic/gin"

// 获取权限列表
func GetRights(c *gin.Context)  {

	rights := map[string]interface{}{
		"code":    0,
		"message": "操作成功",
		"rights": []map[string]interface{}{
			{"id": 1, "right_name": "用户管理", "path": "/users", "level": "一级"},
			{"id": 2, "right_name": "权限管理", "path": "/rights", "level": "二级"},
			{"id": 3, "right_name": "商品管理", "path": "/good", "level": "二级"},
			{"id": 4, "right_name": "订单管理", "path": "/orders", "level": "三级"},
			{"id": 5, "right_name": "数据统计", "path": "/reports", "level": "二级"},
		},
	}

	c.JSON(200, rights)
}
