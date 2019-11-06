package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"gocv.io/x/gocv"
)

func main() {
	// for a good start we just print nice notice about versions
	fmt.Printf("GoCV=%s on OpenCV=%s\n", gocv.Version(), gocv.OpenCVVersion())

	// acquire capture device (camera)
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		log.Fatalf("cannot acquire capture device: %v", err)
	}
	defer webcam.Close()

	// create display windows
	win := gocv.NewWindow("Motion Window")
	win.ResizeWindow(640, 360)
	defer win.Close()

	dbgWin1 := gocv.NewWindow("FOREGROUND")
	dbgWin1.ResizeWindow(640, 360)
	defer dbgWin1.Close()

	dbgWin2 := gocv.NewWindow("BINARY")
	dbgWin2.ResizeWindow(640, 360)
	defer dbgWin2.Close()

	dbgWin3 := gocv.NewWindow("ERODED")
	dbgWin3.ResizeWindow(640, 360)
	defer dbgWin3.Close()

	dbgWin4 := gocv.NewWindow("DILATED")
	dbgWin4.ResizeWindow(640, 360)
	defer dbgWin4.Close()

	img := gocv.NewMat()
	defer img.Close()
	imgDelta := gocv.NewMat()
	defer imgDelta.Close()

	mog2 := gocv.NewBackgroundSubtractorMOG2()
	defer mog2.Close()

	erodeKernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
	defer erodeKernel.Close()
	dilateKernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(15, 15))
	defer dilateKernel.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			continue
		}
		if img.Empty() {
			continue
		}

		// extract foreground
		mog2.Apply(img, &imgDelta)
		dbgWin1.IMShow(imgDelta)

		// convert to binary image
		gocv.Threshold(imgDelta, &imgDelta, 220, 255, gocv.ThresholdBinary)
		dbgWin2.IMShow(imgDelta)

		// Erosion (morphology)
		gocv.Erode(imgDelta, &imgDelta, erodeKernel)
		dbgWin3.IMShow(imgDelta)

		// Dilation (morphology)
		gocv.Dilate(imgDelta, &imgDelta, dilateKernel)
		dbgWin4.IMShow(imgDelta)

		// find contours and surround them with rectangle
		contours := gocv.FindContours(imgDelta, gocv.RetrievalExternal, gocv.ChainApproxSimple)
		for _, c := range contours {
			area := gocv.ContourArea(c)
			if area < 5000 {
				continue
			}
			rect := gocv.BoundingRect(c)
			gocv.Rectangle(&img, rect, color.RGBA{0, 255, 255, 0}, 5)
		}

		// display image enhanced with detected movement regions
		win.IMShow(img)

		// wait 1ms or exit (on Esc key pressed)
		if win.WaitKey(1) == 27 {
			break
		}
	}
}
