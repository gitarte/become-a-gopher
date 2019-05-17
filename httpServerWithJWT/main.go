package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	public := r.Group("/")
	{
		public.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "YOU ARE IN PUBLIC AREA!\n")
		})

		public.POST("/auth", JWTGetToken())
	}

	protected := r.Group("/protected", JWTCheckToken())
	{
		protected.GET("/admin", func(c *gin.Context) {
			c.String(http.StatusOK, "YOU ARE IN RESTRICTED AREA!\n")
		})
	}

	r.Run()
}
