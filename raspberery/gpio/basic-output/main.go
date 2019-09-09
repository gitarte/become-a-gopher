package main

import (
	"fmt"
	"os"
	"time"

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

	// "do forever" as you usually do in embedded programming
	// first the trivial case for presentation purpose only
	// I need it just to remember to provide an introduction of some concepts
	// for {
	// 	pin1.Toggle()
	// 	pin2.Toggle()
	// 	time.Sleep(1 * time.Second)
	// }

	// "do forever" as you usually do in embedded programming
	// now the better solution
	tmr1 := time.NewTimer(1 * time.Second)
	tmr2 := time.NewTimer(3 * time.Second)
	for {
		select {
		case <-tmr1.C:
			pin1.Toggle()
			fmt.Println("PIN 19 chage")
		case <-tmr2.C:
			pin2.Toggle()
			fmt.Println("PIN 21 chage")
		}
	}
}
