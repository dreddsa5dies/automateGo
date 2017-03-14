package main

import (
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/tealeg/xlsx"
)

var opts struct {
	FileXLSX string `short:"s" long:"save" default:"test.xlsx" description:"Региональная численность населения"`
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

	// переменные
	// файл
	var file *xlsx.File
	// страница
	var sheet *xlsx.Sheet
	// строка
	var row *xlsx.Row
	// ячейка
	var cell *xlsx.Cell

	// создаем новый файл
	file = xlsx.NewFile()
	// добавляем страницу
	sheet, err = file.AddSheet("Sheet")
	log.Printf("Добавление страницы %v", sheet)
	if err != nil {
		log.Printf("Ошибка добаления страницы %v", err)
	}

	for i := 0; i < 50; i++ {
		// создаем овую строку в странице
		row = sheet.AddRow()
		log.Printf("Добавление строки %v", row)
		// добавляем ячейки в строку (следуют одна за одной)
		row.AddCell().SetString("Col A")
		row.AddCell().SetString("Col B")
		row.AddCell().SetString("Col C")
		// и еще добавлем
		cell = row.AddCell()
		log.Printf("Добавление ячейки %v", cell)
		cell.Value = "Spam Spam Spam"
		// в цикле создается 50 строк
	}
	// сохранение
	err = file.Save(pwdDir + "/" + opts.FileXLSX)
	if err != nil {
		log.Printf("Ошибка сохранения файла %v", err)
	}
	log.Printf("Сохранение в %v", opts.FileXLSX)

	// готово
	log.Println("Готово")
}
