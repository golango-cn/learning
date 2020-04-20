package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func Upload(c *gin.Context) {

	file, err := c.FormFile("file")

	pwd, _ := os.Getwd()
	dest := pwd + "/upload/"
	os.MkdirAll(dest, os.ModePerm)

	dest += file.Filename
	err = c.SaveUploadedFile(file, dest)

	fmt.Println(err)

	if err != nil {
		c.JSON(200, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, map[string]interface{}{
			"code": 200,
		})
	}

}
