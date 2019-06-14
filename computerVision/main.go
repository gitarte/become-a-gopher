package main

//git clone  https://github.com/opencv/opencv.git

import (
	"fmt"
	"image/color"
	"time"

	"gocv.io/x/gocv"
)

func justShowCamra() {
	camera, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("sia la la la")
	img := gocv.NewMat()
	for {
		camera.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	xmlFile := "/home/artgaw/Documents/opencv/data/haarcascades/haarcascade_frontalface_default.xml"

	// load camera
	camera, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer camera.Close()

	// load display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	for {
		if ok := camera.Read(&img); !ok {
			fmt.Printf("cannot read device 0\n")
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)

		// draw a rectangle around each face on the original image,
		// along with text identifying as "Human"
		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)
		}

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}

		time.Sleep(200 * time.Millisecond)
	}
}
