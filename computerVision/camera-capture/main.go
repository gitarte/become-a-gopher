package main

import (
	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.VideoCaptureDevice(2)
	window := gocv.NewWindow("JAK MI WYSTAWICIE POCHLEBNE OCENY TO NA NASTĘPNYM MEETUPIE POKAŻĘ WAM, JAK SIĘ ROBI TAKIE OKNO")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
