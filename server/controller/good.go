package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ndb/server/model"
	"strconv"
)

// 商品分类
func GetCates(c *gin.Context) {

	cates := map[string]interface{}{
		"code":       0,
		"message":    "操作成功",
		"totalCount": 200,
		"cates": []map[string]interface{}{
			{
				"id":          1,
				"cate_name":   "大家电",
				"cate_delete": false,
				"cate_level":  0,
				"children": []map[string]interface{}{
					{
						"id":          101,
						"cate_name":   "电视",
						"cate_delete": true,
						"cate_level":  1,
						"children": []map[string]interface{}{
							{
								"id":          1011,
								"cate_name":   "TCL",
								"cate_delete": false,
								"cate_level":  2,
							},
							{
								"id":          1012,
								"cate_name":   "联想",
								"cate_delete": true,
								"cate_level":  2,
							},
						},
					},
					{
						"id":          201,
						"cate_name":   "手机",
						"cate_delete": false,
						"cate_level":  1,
						"children": []map[string]interface{}{
							{
								"id":          2011,
								"cate_name":   "三星",
								"cate_delete": true,
								"cate_level":  2,
							},
							{
								"id":          2012,
								"cate_name":   "华为",
								"cate_delete": false,
								"cate_level":  2,
							},
						},
					},
				}},
			{"id": 3, "cate_name": "电冰箱", "cate_level": 0},
			{"id": 4, "cate_name": "洗衣机", "cate_level": 0},
		},
	}

	c.JSON(200, cates)

}

// 属性参数
func GetAttrs(c *gin.Context) {

	attr_type := c.Param("attr_type")

	var json map[string]interface{}
	c.ShouldBind(&json)

	attr_type_name := map[string]string{
		"first": "动态参数", "second": "静态属性",
	}
	var mps []map[string]interface{}
	type_name := attr_type_name[attr_type]

	for i := 10; i > 0; i-- {

		mp := make(map[string]interface{})
		mp["id"] = i
		mp["name"] = fmt.Sprintf("%s%d", type_name, i)

		var tags []string
		for j := 0; j < 5; j++ {
			tags = append(tags, fmt.Sprintf("标签%d_%d", i, j))
		}
		mp["tags"] = tags

		mps = append(mps, mp)
	}

	c.JSON(200, mps)

}

// 商品列表
func GetGoods(c *gin.Context) {

	pageIndex, _ := strconv.ParseInt(c.DefaultQuery("pageIndex", "0"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.DefaultQuery("pageSize", "10"), 10, 64)

	var query map[string]interface{}
	c.ShouldBind(&query)

	var keys string
	if v, has := query["keys"]; has {
		keys = v.(string)
	}


	var g model.Good
	goods, total, _ := g.Gets(keys, int(pageIndex), int(pageSize))

	results := map[string]interface{}{
		"code":       0,
		"message":    "操作成功",
		"totalCount": total,
		"goods":      goods,
	}

	c.JSON(200, results)
}
