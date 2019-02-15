package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// File -
type File struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
}

// Result -
type Result struct {
	Result []File `json:"result"`
}

// List -
func List(c *gin.Context) {
	files, err := ioutil.ReadDir("/files/")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		log.Error().Msg(err.Error())
		return
	}

	result := make([]File, 1, 100)
	for i, f := range files {
		result = append(result, File{i, f.Name()})
	}

	c.JSON(http.StatusOK, result)
}

// Read -
func Read(c *gin.Context) {
	file := c.Param("file")
	data, err := ioutil.ReadFile("/files/" + file)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, string(data))
}
