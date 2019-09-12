package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	duration1 = 1 * time.Second
	duration2 = 3 * time.Second
)

func main() {
	// open an access to GPIO
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		rpio.Close()
		fmt.Println("GPIO access has been closed")
	}()

	// accuire the objects that resembles two pins of GPIO
	pin1 := rpio.Pin(10) //	BCM pin 10 is physical pin 19
	pin2 := rpio.Pin(9)  //	BCM pin 09 is physical pin 21

	// set pins as outputs
	pin1.Output()
	pin2.Output()

	// set initial states
	pin1.High() // set 3.3V on pin 19
	pin2.Low()  // set 0.0V on pin 21

	// create two timers with duration of 1 and 3 seconds
	tmr1 := time.NewTimer(duration1)
	tmr2 := time.NewTimer(duration2)

	// we need to be able to handle Ctrl+C gracefully
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	// do forever
	// we implement infinite loop as you usually do in embedded programming
	for {
		select {
		case <-tmr1.C:
			pin1.Toggle()                // change pin 19 to it's opposite state
			tmr1.Reset(duration1)        // we have used timer 1, so we must reset it before next tick
			fmt.Println("toggle pin 19") //
		case <-tmr2.C:
			pin2.Toggle()                // change pin 21 to it's opposite state
			tmr2.Reset(duration2)        // we have used timer 2, so we must reset it before next tick
			fmt.Println("toggle pin 21") //
		case <-sig:
			pin1.Low()               // switch off the pin 19
			pin2.Low()               // switch off the pin 21
			fmt.Println("interrupt") //
			return
		}
	}
}
