// image
package main

import (
	"fmt"
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

	log.Printf("Размеры: %v", bounds.String())
	log.Printf("Размеры: %v", bounds.Size().String())

	// установить новое изображение в размерах от 1/4 старого
	m := image.NewNRGBA(image.Rect(pointCheckX/4, pointCheckY/4, pointCheckX/4+pointCheckX/2, pointCheckY/4+pointCheckY/2))
	// создать это изображение (обрезав)
	draw.Draw(m, m.Bounds(), img, image.Point{pointCheckX / 4, pointCheckY / 4}, draw.Src)

	// файл для сохранения
	toimg, _ := os.Create("./new.png")
	defer toimg.Close()

	// записать в формате png
	png.Encode(toimg, m)
	log.Printf("Записано в: %v", toimg.Name())

	// cropped.png
	n := image.NewNRGBA(image.Rect(335, 345, 565, 560))
	draw.Draw(n, n.Bounds(), img, image.Point{335, 345}, draw.Src)
	toN, _ := os.Create("./cropped.png")
	defer toN.Close()
	png.Encode(toN, n)
	log.Printf("Записано в: %v", toN.Name())

	// копирование и вставка (copy/paste)
	// открытие файлов
	imgFile1, err := os.Open("./zophie.1.png")
	imgFile2, err := os.Open("./cropped.1.png")
	if err != nil {
		fmt.Println(err)
	}
	defer imgFile1.Close()
	defer imgFile2.Close()
	// декодирование изображения
	img1, _, err := image.Decode(imgFile1)
	img2, _, err := image.Decode(imgFile2)
	if err != nil {
		log.Println(err)
	}
	// нахождение позиции в 2м изображении (у изображения будет высота как у 2го изображения)
	sp2 := image.Point{img1.Bounds().Dx(), 0}
	// новый rectangle (прямоугольник) во 2м изображении
	r2 := image.Rectangle{sp2, sp2.Add(img2.Bounds().Size())}
	// прямойгольник большого изображения
	r := image.Rectangle{image.Point{0, 0}, r2.Max}
	rgba := image.NewRGBA(r)
	// отрисовка 2х изображений рядом
	draw.Draw(rgba, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, img2, image.Point{0, 0}, draw.Src)
	// создание файла созранения
	out, err := os.Create("./pasted.png")
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	// запись
	png.Encode(out, rgba)
	log.Printf("Записано в: %v", out.Name())

	// заполнение вырезанным изображением
	catImg, err := os.Open("./zophie.png")
	if err != nil {
		log.Fatalf("Ошибка открытия файла %v", err)
	}
	defer catImg.Close()
	log.Printf("Открытие файла: %v", catImg.Name())
	catImage, _, err := image.Decode(catImg)
	if err != nil {
		log.Panicf("Ошибка декодирования изображения %v", err)
	}
	catImWidth := catImage.Bounds().Size().X
	catImHeight := catImage.Bounds().Size().Y
	faceIm := image.NewNRGBA(image.Rect(0, 0, catImWidth, catImHeight))
	for left := 0; left < catImWidth; left += catImWidth / 2 {
		for top := 0; top < catImHeight; top += catImHeight / 2 {
			println(left, top)
			draw.Draw(faceIm, faceIm.Bounds(), catImage, image.Point{left, top}, draw.Src)
		}
	}
	topng, _ := os.Create("./copyTo.png")
	defer topng.Close()
	png.Encode(topng, faceIm)
	log.Printf("Записано в: %v", topng.Name())
}
