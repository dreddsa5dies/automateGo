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
	// ткрыть файл
	file, _ := os.OpenFile("./zophie.png", os.O_RDONLY, 0600)
	defer file.Close()
	log.Printf("Открытие файла: %v", file.Name())

	// декодировать его как изображение
	img, _, err := image.Decode(file)
	if err != nil {
		log.Panicf("Ошибка декодирования изображения %v", err)
	}

	bounds := img.Bounds()
	// всего X, Y (количество)
	pointCheckX := bounds.Max.X
	pointCheckY := bounds.Max.Y
	log.Printf("Всего X: %v", pointCheckX)
	log.Printf("Всего X: %v", pointCheckY)

	// установить новое изображение в размерах от 1/4 старого
	m := image.NewNRGBA(image.Rect(pointCheckX/4, pointCheckY/4, pointCheckX/4+pointCheckX/2, pointCheckY/4+pointCheckY/2))
	// создать это изображение (обрезав)
	draw.Draw(m, m.Bounds(), img, image.Point{pointCheckX / 4, pointCheckY / 4}, draw.Src)

	// файл для сохранения
	toimg, _ := os.Create("./new.png")
	defer toimg.Close()

	// записать в формате png
	png.Encode(toimg, m)
}
