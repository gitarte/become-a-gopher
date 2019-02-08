package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//	we configure global logger
	LoadZerolog()

	//	we setup GIR router
	r := gin.New()
	//	this captures any panic and lets the app live
	r.Use(gin.Recovery())
	//	we pass our custom logger
	r.Use(Logger())

	//	we define route that is bublic
	root := r.Group("/")
	{
		root.GET("/list", List)
		root.GET("/read/:file", Read)
	}

	//	we define route that is restricted
	admin := r.Group("/admin")
	{
		admin.POST("/write/:file", Write)
	}

	r.Run(":8080")
}
