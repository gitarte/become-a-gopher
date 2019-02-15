package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//	we configure global logger
	LoadZerolog()

	//	configure GIN
	r := gin.New()                                // setup GIN router
	r.Use(gin.Recovery())                         // this captures any panic and lets the app live
	r.Use(Logger())                               // pass custom logger
	r.Static("/resources", "./resources")         // indicate the source of static content
	r.StaticFile("/favicon.ico", "./favicon.ico") // set custom favicon
	r.LoadHTMLGlob("templates/*")                 // set the path to template files

	//	we define route that is bublic
	root := r.Group("/")
	{
		root.GET("/list", List)
		root.GET("/read/:file", Read)
	}

	//	we define route that is restricted
	admin := r.Group("/admin", gin.BasicAuth(AdminUsers))
	{
		admin.GET("", AdminRoot)
		admin.GET("/read/:file", AdminRead)
		admin.POST("/write", AdminWrite)
	}

	r.Run(":8080")
}
