// тестовое github.com/anthonynsimon/bild
package main

import (
	"image"
	"log"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

func main() {
	// открытие файла
	img, err := imgio.Open("zophie.png")
	if err != nil {
		log.Fatalln(err)
	}

	// обрезка
	crop := transform.Crop(img, image.Rect(335, 345, 565, 560))
	println(crop.Bounds().Size().X, crop.Bounds().Size().Y)

	// сохранение (имя, что, в каком формате)
	if err := imgio.Save("crop", crop, imgio.PNG); err != nil {
		log.Fatalln(err)
	}

	// изменение размера (уменьшение в 4 раза)
	resize := transform.Resize(img, img.Bounds().Size().X/4, img.Bounds().Size().Y/4, transform.Linear)
	if err := imgio.Save("resize", resize, imgio.PNG); err != nil {
		log.Fatalln(err)
	}
	println(resize.Bounds().Size().X, resize.Bounds().Size().Y)

	// TODO: вставка
}
