package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// cli
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %v Data\n", os.Args[0])
		os.Exit(1)
	} else {
		// data for barcode
		var strData string
		for k, v := range os.Args[1:] {
			if k != len(os.Args[1:])-1 {
				strData += v + " "
			} else {
				strData += v
			}
		}
		// Create the barcode
		qrCode, _ := qr.Encode(strData, qr.M, qr.Auto)

		// Scale the barcode to 200x200 pixels
		qrCode, _ = barcode.Scale(qrCode, 200, 200)

		// create the output file
		file, _ := os.Create("qrcode.png")
		defer file.Close()

		// encode the barcode as png
		png.Encode(file, qrCode)
	}
}
