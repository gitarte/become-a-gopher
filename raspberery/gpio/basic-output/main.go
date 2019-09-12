package main

import (
	"fmt"
	"os"
	"sync"
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

	// create two timers with duration of 1 and 3 seconds
	const duration1 = 1 * time.Second
	const duration2 = 3 * time.Second
	tmr1 := time.NewTimer(duration1)
	tmr2 := time.NewTimer(duration2)

	// do forever
	// we implement infinite loop as you usually do in embedded programming
	// except here we use timers that send the event through channels so we need goroutine
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-tmr1.C:
				pin1.Toggle()         // change pin 19 to it's opposite state
				tmr1.Reset(duration1) // we have used timer 1, so we must reset it before next tick
				fmt.Println("PIN 19 chage")
			case <-tmr2.C:
				pin2.Toggle()         // change pin 21 to it's opposite state
				tmr2.Reset(duration2) // we have used timer 2, so we must reset it before next tick
				fmt.Println("PIN 21 chage")
			}
		}
	}()
	wg.Wait()
}
