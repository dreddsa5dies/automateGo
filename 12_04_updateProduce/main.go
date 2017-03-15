package main

import (
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/tealeg/xlsx"
)

var opts struct {
	FileXLSX string `short:"o" long:"open" default:"produceSales.xlsx" description:"Файл для работы"`
}

func main() {
	// разбор флагов
	flags.Parse(&opts)

	// в какой папке исполняемы файл
	pwdDir, _ := os.Getwd()

	// создание файла log для записи ошибок
	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	// запись ошибок и инфы в файл
	log.SetOutput(fLog)

	excelFileName := pwdDir + "/" + opts.FileXLSX

	log.Printf("Открытие рабочей книги: %v", opts.FileXLSX)

	wb, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("Ошибка открытия книги %v", err)
	}

	sheet := wb.Sheet["Sheet"]

	// Типы продукции и их обновленные цены
	priceUpdates := map[string]float64{
		"Lemon":  3.07,
		"Celery": 1.19,
		"Garlic": 1.27,
	}

	//TODO: создать цикл по сторкам и обновить цены
	// в github.com/tealeg/xlsx все индексы и номера начинаются с 0
	for rowNum := 1; rowNum < sheet.MaxRow; rowNum++ {
		produceName := sheet.Cell(rowNum, 0).Value
		if _, ok := priceUpdates[produceName]; ok {
			sheet.Cell(rowNum, 1).SetFloat(priceUpdates[produceName])
			log.Printf("Замена в строке %v начения для %v", rowNum, produceName)
		}
	}

	wb.Save("update" + opts.FileXLSX)

	// готово
	log.Println("Готово")
}
