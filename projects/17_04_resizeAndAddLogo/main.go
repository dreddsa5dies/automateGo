// вставка логотипа
package main

import (
	"io"
	"log"
	"os"
	"strings"

	"image"

	"github.com/disintegration/imaging"
	flags "github.com/jessevdk/go-flags"
)

const (
	squareFitSize = 300
	logoFilename  = "catlogo.png"
)

var opts struct {
	SRC string `short:"o" long:"open" default:"./src" description:"Исходная папка"`
	DST string `short:"s" long:"save" default:"./dst" description:"Директория для сохранения"`
}

func main() {
	// разбор флагов
	flags.Parse(&opts)

	// создание файла log для записи ошибок
	fLog, err := os.OpenFile(`./.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	// запись ошибок и инфы в файл и вывод
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))

	// TODO:  организовать цикл по всем файлам в папке
	workFunc(opts.SRC)
}

func workFunc(dir string) {
	// открыти файла с логотипом
	logoIm, err := imaging.Open(logoFilename)
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}

	// открываем директорию
	dh, err := os.Open(dir)
	if err != nil {
		return
	}
	defer dh.Close()
	// считывание списка файлов
	for {
		fis, err := dh.Readdir(10)
		if err == io.EOF {
			break
		}
		// проход по файлам
		for _, fi := range fis {
			if strings.HasSuffix(fi.Name(), ".png") {
				// TODO: проверить, нуждается ли изображение в изменении размеров
				// открытие файла изображения
				log.Printf("Открытие %v", fi.Name())
				img, err := imaging.Open(opts.SRC + "/" + fi.Name())
				if err != nil {
					log.Fatalf("Ошибка открытия файла: %v", err)
				}
				// проверка необходимости изменения размеров
				wigth := img.Bounds().Size().X
				height := img.Bounds().Size().Y
				if wigth > squareFitSize || height > squareFitSize {
					// расчет новых значений
					if wigth > height {
						height = int((squareFitSize / wigth) * height)
						wigth = squareFitSize
					} else {
						wigth = int((squareFitSize / height) * wigth)
						height = squareFitSize
					}
					// изменение размеров изображения
					log.Printf("Изменение размеров изображения %v", fi.Name())
					img = imaging.Resize(img, wigth, height, imaging.Lanczos)
				}

				// какой-то большой логотип
				logoImN := imaging.Resize(logoIm, wigth/2, height/2, imaging.Lanczos)
				logoWidth := logoImN.Bounds().Size().X
				logoHeight := logoImN.Bounds().Size().Y

				// добавить логотип
				log.Printf("Добавление логотипа в %v", fi.Name())
				img = imaging.Overlay(img, logoImN, image.Pt(wigth-logoWidth, height-logoHeight), 0.5)

				// сохранить
				err = imaging.Save(img, opts.DST+"/"+fi.Name())
				if err != nil {
					log.Fatalf("Save failed: %v", err)
				}

			}
			// рекурсивный проход по поддиректориям
			if fi.IsDir() {
				workFunc(dir + "/" + fi.Name())
			}
		}
	}
}
