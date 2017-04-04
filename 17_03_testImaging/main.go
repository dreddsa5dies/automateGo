// тестовое disintegration/imaging
package main

import (
	"image"
	"image/color"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	// открытие файла
	src, err := imaging.Open("./zophie.png")
	if err != nil {
		log.Fatalf("Open failed: %v", err)
	}

	// Обрежьте исходное изображение размером 350x350 пикселей с помощью центрального якоря.
	src = imaging.CropAnchor(src, 350, 350, imaging.Center)

	// Измените размер обрезанного изображения до width = 256px, сохраняя пропорции.
	src = imaging.Resize(src, 256, 0, imaging.Lanczos)

	// Создайте размытую версию изображения.
	img1 := imaging.Blur(src, 2)

	// Создайте полутоновую версию изображения с более высоким контрастом и резкостью.
	img2 := imaging.Grayscale(src)
	img2 = imaging.AdjustContrast(img2, 20)
	img2 = imaging.Sharpen(img2, 2)

	// Создайте инвертированную версию изображения.
	img3 := imaging.Invert(src)

	// Создайте тисненый вариант изображения, используя фильтр свертки.
	img4 := imaging.Convolve3x3(
		src,
		[9]float64{
			-1, -1, 0,
			-1, 1, 1,
			0, 1, 1,
		},
		nil,
	)

	// Создайте новое изображение и вставьте в него четыре созданных изображения.
	dst := imaging.New(512, 512, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img1, image.Pt(0, 0))
	dst = imaging.Paste(dst, img2, image.Pt(0, 256))
	dst = imaging.Paste(dst, img3, image.Pt(256, 0))
	dst = imaging.Paste(dst, img4, image.Pt(256, 256))

	// Сохраните полученное изображение в формате PNG.
	err = imaging.Save(dst, "./test.png")
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}

	// обрезка изображения
	file, err := imaging.Open("./zophie.png")
	if err != nil {
		log.Fatalf("Open failed: %v", err)
	}
	crop := imaging.Crop(file, image.Rect(335, 345, 565, 560))
	err = imaging.Save(crop, "./crop.png")
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}

	// изменение размера (уменьшение в 4 раза) сохраняя пропорции
	resize := imaging.Resize(file, file.Bounds().Size().X/4, 0, imaging.Lanczos)
	err = imaging.Save(resize, "./resize.png")
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}

	// вставка в центре
	paste := imaging.PasteCenter(file, resize)
	err = imaging.Save(paste, "./paste.png")
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}

	// вставка куда-нибудь
	pasteN := imaging.Paste(file, resize, image.Pt(0, 0))
	err = imaging.Save(pasteN, "./pasteN.png")
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}

	// повороты
	rotate := imaging.Rotate90(file)
	err = imaging.Save(rotate, "./rotate90.png")
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}

	// проход по pixel
	pixels := imaging.New(100, 100, color.NRGBA{0, 0, 0, 0})
	for i := 0; i < 100; i++ {
		for j := 0; j < 50; j++ {
			pixels.SetNRGBA(i, j, color.NRGBA{0, 100, 100, 255})
		}
		for j := 50; j < 100; j++ {
			pixels.SetNRGBA(i, j, color.NRGBA{100, 100, 0, 100})
		}
	}
	err = imaging.Save(pixels, "./pixels.png")
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}
}
