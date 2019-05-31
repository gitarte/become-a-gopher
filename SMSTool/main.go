package main

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// Message -
type Message struct {
	Content string `json:"content"`
	Number  string `json:"number"`
}

func main() {
	r := gin.Default()

	r.POST("/sms", func(c *gin.Context) {
		m := Message{}
		err := c.Bind(&m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"bind status": err.Error()})
			return
		}

		cmd := exec.Command("./write.sh", m.Content, m.Number)
		err = cmd.Run()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"send status": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "poszło"})
		return
	})

	r.Run(":8080")
}
