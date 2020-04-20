package controller

import "github.com/gin-gonic/gin"

// 获取菜单列表
func GetMenus(c *gin.Context) {

	menus := map[string]interface{}{
		"code":    0,
		"message": "操作成功",
		"menus": []map[string]interface{}{
			{"id": 1, "name": "用户管理", "path": "/users", "children": []map[string]interface{}{
				{"id": 1100, "name": "用户列表", "path": "/users"},
			}},
			{"id": 2, "name": "权限管理", "path": "/rights", "children": []map[string]interface{}{
				{"id": 2100, "name": "角色列表", "path": "/roles"},
				{"id": 2101, "name": "权限列表", "path": "/rights"},
			}},
			{"id": 3, "name": "商品管理", "path": "/good", "children":[]map[string]interface{}{
				{"id": 3100, "name": "商品分类", "path": "/cates"},
				{"id": 3101, "name": "分类参数", "path": "/params"},
				{"id": 3102, "name": "商品列表", "path": "/goods"},
			}},
			{"id": 4, "name": "订单管理", "path": "/orders"},
			{"id": 5, "name": "数据统计", "path": "/reports"},

			{"id": 6, "name": "通用", "path": "/common", "children":[]map[string]interface{}{
				{"id": 6100, "name": "通用", "path": "/common"},
			}},
		},
	}

	c.JSON(200, menus)

}
