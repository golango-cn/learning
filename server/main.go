package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"ndb/server/router"
)

func main()  {

	gin.SetMode(gin.DebugMode)
	router := router.InitRouter()

	if err := router.Run(":8092"); err != nil {
		log.Fatal(err)
	}

}
