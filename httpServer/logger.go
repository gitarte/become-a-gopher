package main

import (
	"time"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LoadZerolog - this function configures globaly available instance of zerolog
func LoadZerolog() {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.000"
	log.Logger = zerolog.
		New(os.Stdout).
		With().
		Timestamp().
		Logger()
}

// Logger - the logger and it's parameters
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//	before handler
		start := getNowInMillis()

		//	do the handler
		c.Next()

		// 	after handler
		end := getNowInMillis()
		status := c.Writer.Status()

		var event *zerolog.Event
		if status >= 100 && status < 400 {
			event = log.Info()
		} else {
			event = log.Error()
		}

		event.
			Str("protocol", c.Request.Proto).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int64("start", start).
			Int64("end", end).
			Int64("duration", end-start).
			Str("from", c.ClientIP()).
			Str("User-Agent", c.Request.Header.Get("User-Agent")).
			Str("X-Request-Id", c.Request.Header.Get("X-Request-ID")).
			Str("X-Forwarded-For", c.Request.Header.Get("X-Forwarded-For")).
			Int("status", status).
			Msg(c.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}

func getNowInMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
