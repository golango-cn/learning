package controller

import "github.com/gin-gonic/gin"

func GetRules(c *gin.Context)  {
	rules := []map[string]interface{}{
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
			"id":        202,
			"auth_name": "订单管理",
			"path":      "orders",
			"children": []map[string]interface{}{
				{
					"id":        2011,
					"auth_name": "商品修改",
					"path":      "orders",
					"children": []map[string]interface{}{
						{
							"id":        20111,
							"auth_name": "修改",
							"path":      "orders",
						},
						{
							"id":        20112,
							"auth_name": "报表",
							"path":      "orders",
						},
						{
							"id":        20113,
							"auth_name": "删除",
							"path":      "orders",
						},
					},
				},
				{
					"id":        2012,
					"auth_name": "订单删除",
					"path":      "orders",
				},
			},
		},
	}
	c.JSON(200, rules)
}
