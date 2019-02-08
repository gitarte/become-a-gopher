package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Write -
func Write(c *gin.Context) {
	log.Info().Str("lalala", "alalala").Msg("")
}
