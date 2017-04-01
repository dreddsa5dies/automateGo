// image
package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	file, _ := os.OpenFile("./zophie.png", os.O_RDONLY, 0600)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Panicf("Ошибка декодирования изображения %v", err)
	}

	m := image.NewNRGBA(image.Rect(0, 0, 800, 600))
	draw.Draw(m, m.Bounds(), img, image.Point{0, 0}, draw.Src)

	toimg, _ := os.Create("./new.png")
	defer toimg.Close()

	png.Encode(toimg, m)
}
