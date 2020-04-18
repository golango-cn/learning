package controller

import "github.com/gin-gonic/gin"

func GetRoles(c *gin.Context) {

	roles := map[string]interface{}{
		"code":    0,
		"message": "操作成功",
		"roles": []map[string]interface{}{
			{
				"id":        1,
				"role_name": "超级管理员",
				"role_desc": "非常厉害的用户",
				"children": []map[string]interface{}{
					{
						"id":        101,
						"auth_name": "商品管理",
						"path":      "good",
						"children": []map[string]interface{}{
							{
								"id":        1011,
								"auth_name": "商品修改",
								"path":      "good",
								"children": []map[string]interface{}{
									{
										"id":        10111,
										"auth_name": "修改",
										"path":      "good",
									},
									{
										"id":        10112,
										"auth_name": "报表",
										"path":      "good",
									},
									{
										"id":        10113,
										"auth_name": "删除",
										"path":      "good",
									},
								},
							},
							{
								"id":        1012,
								"auth_name": "商品删除",
								"path":      "good",
							},
						},
					},
					{
						"id":        201,
						"auth_name": "订单管理",
						"path":      "good",
						"children": []map[string]interface{}{
							{
								"id":        2011,
								"auth_name": "商品修改",
								"path":      "good",
							},
							{
								"id":        2012,
								"auth_name": "商品删除",
								"path":      "good",
							},
						},
					},
				}},
			{"id": 3, "role_name": "管理员", "role_desc": ""},
			{"id": 4, "role_name": "普通用户", "role_desc": ""},
		},
	}

	c.JSON(200, roles)

}
