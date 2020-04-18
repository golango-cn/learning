package controller

import "github.com/gin-gonic/gin"

// 获取用户列表
func GetUsers(c *gin.Context) {

	users := map[string]interface{}{
		"code":    0,
		"message": "操作成功",
		"users": []map[string]interface{}{
			{"id": 1, "role_name": "超级管理员", "user_name": "admin", "create_time": 1486720211, "mobile": "152****1111", "status": true},
			{"id": 3, "role_name": "超级管理员", "user_name": "admin", "create_time": 1486720211, "mobile": "152****1111", "status": false},
			{"id": 4, "role_name": "超级管理员", "user_name": "admin", "create_time": 1486720211, "mobile": "152****1111", "status": true},
		},
		"totalCount": 30,
	}

	c.JSON(200, users)

}

// 登录
func Login(c *gin.Context) {

	var json map[string]interface{}
	c.ShouldBind(&json)

	name := json["name"]
	pwd := json["pwd"]

	if name == "admin" && pwd == "123456" {
		c.JSON(200, map[string]interface{}{
			"code":    0,
			"message": "登录成功",
		})
	} else {
		c.JSON(200, map[string]interface{}{
			"code":    -100,
			"message": "用户名或密码错误",
		})
	}

}
