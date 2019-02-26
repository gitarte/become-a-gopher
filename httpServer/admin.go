package main

import (
	"database/sql"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// AdminUsers -
var AdminUsers = gin.Accounts{
	"foo": "bar",
	"du":  "pa",
}

// Form -
type Form struct {
	File    string `form:"file"    binding:"required"`
	Content string `form:"content" binding:"required"`
}

// AdminRoot -
func AdminRoot(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"title":     "Admin root",
			"where":     "Hello admin",
			"condition": false,
		})
	}
}

// AdminRead -
func AdminRead(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.Param("file")
		data, err := ioutil.ReadFile("/files/" + file)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}

		c.HTML(http.StatusOK, "adminWrite.html", gin.H{
			"title":     "Admin write",
			"where":     file,
			"file":      file,
			"content":   string(data),
			"condition": false,
		})
	}
}

// AdminWrite -
func AdminWrite(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form Form

		err := c.ShouldBind(&form)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		log.Debug().Msgf("%+v", form)

		err = ioutil.WriteFile("/files/"+form.File, []byte(form.Content), 0644)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(http.StatusOK, "OK")
	}
}
