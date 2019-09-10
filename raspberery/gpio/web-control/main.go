package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	// open an access to GPIO
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	// accuire the objects that resembles two pins of GPIO
	pin1 := rpio.Pin(10) //	BCM pin 10 is physical pin 19
	pin2 := rpio.Pin(9)  //	BCM pin 09 is physical pin 21

	// set pins as outputs
	pin1.Output()
	pin2.Output()

	// set initial states
	pin1.High() // set 3.3V on pin 19
	pin2.Low()  // set 0.0V on pin 21

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// For each matched request Context will hold the route definition
	router.POST("/toggle/:pin", func(p1, p2 rpio.Pin) gin.HandlerFunc {
		return func(c *gin.Context) {
			pin, err := strconv.ParseInt(c.Param("pin"), 10, 8)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
			}

			switch pin {
			case 1:
				p1.Toggle()
				c.String(http.StatusOK, "pin1 oK")
			case 2:
				p2.Toggle()
				c.String(http.StatusOK, "pin2 oK")
			default:
				c.String(http.StatusBadRequest, "wrong pin number")
			}
		}
	}(pin1, pin2))

	router.Run(":8080")
}
