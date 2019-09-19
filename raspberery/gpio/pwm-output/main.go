package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/stianeikeland/go-rpio"
)

const (
	bcmPin         = 19              // BCM pin 19 is physical pin 35
	servoMostLeft  = 32              // 0.5ms pulse, duty cycle 2.5%
	servoCenter    = 96              // 1.5ms pulse, duty cycle 7.5%
	servoMostRight = 160             // 2.5ms pulse, duty cycle 12.5%
	clockFrequency = 64000           // internal clock that controlls PWM will oscilate with frequency of 64kHz
	cycleWidth     = 1280            // PWM will have freqency of 64000/12800 = 50Hz
	interval       = 1 * time.Second // we will move the shaft of the servo in every single second
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

	// accuire the objects that resembles the pin of GPIO
	pin := rpio.Pin(bcmPin)  // BCM pin 19 is physical pin 35
	pin.Mode(rpio.Pwm)       // this pin is able to generate hardware PWM signal
	pin.Freq(clockFrequency) // internal clock that controlls PWM on this pin will oscilate with frequency oh 64kHz

	// create a timer with duration of 1 second
	tmr := time.NewTimer(interval)

	// we need to be able to handle Ctrl+C gracefully
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	// prepare a slice of available servo positions
	servoPositions := []uint32{
		servoMostLeft,
		servoCenter,
		servoMostRight,
	}

	i := 0                   // this is just iteration counter
	l := len(servoPositions) // we need a number of servo positions that we defined
	for {
		select {
		case <-tmr.C:
			// here we select next position of servo
			pos := i % l // servo position that we want to set depends on current iteration count divided modulo by the number of available positions
			i++          // increase the counter for next iteration
			
			// here we set choosen position
			pin.DutyCycle(servoPositions[pos], cycleWidth) // we set the length of the pulse
			
			// Beware here!
			// Using timers without Reset inside infinite loop will cause a DEADLOCK.
			// Therefore it might be considered safer to use NewTicker instead NewTimer
			// Or just do not forget to reset the timer
			tmr.Reset(interval)                            // we have used timer 1, so we must reset it before next tick
			fmt.Printf("servo set in position %d\n", pos)
		case <-sig: // Ctrl+C captured
			// we set servo in center position and leave the program
			pin.DutyCycle(servoPositions[1], cycleWidth) 
			fmt.Println("interupt")
			fmt.Println("servo set in center position")
			return
		}
	}
}
