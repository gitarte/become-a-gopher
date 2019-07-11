package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	go interval(60*time.Second, func(time.Time) {
		fmt.Println("DUPA")
	})

	app := gin.Default()
	app.Use(favicon.New("/app/favicon.ico"))
	app.GET("/ping", func(c *gin.Context) {
		t := time.Now()
		c.String(http.StatusOK, "PONG "+t.Format(time.RFC3339))
	})
	app.Run(":8080")
}
