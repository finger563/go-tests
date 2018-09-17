package main

// #cgo CPPFLAGS: -I../../../Downloads/dlib-19.15
// #cgo LDFLAGS: -ldlib -lc++
// #include "dlib/image_processing/frontal_face_detector.h"
import "C"

import (
	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		if (window.WaitKey(1) == 27) {
			break
		}
	}
}
