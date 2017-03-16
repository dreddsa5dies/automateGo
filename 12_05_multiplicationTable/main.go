package main

import (
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/tealeg/xlsx"
)

var opts struct {
	N        int    `short:"n" long:"number" default:"6" description:"Число N для таблицы"`
	FileSAVE string `short:"s" long:"save" default:"table.xlsx" description:"Файл таблицы"`
}

func main() {
	// разбор флагов
	flags.Parse(&opts)

	// в какой папке исполняемый файл
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

	// создаем новый файл
	file = xlsx.NewFile()
	// добавляем страницу
	sheet, err = file.AddSheet("Sheet")
	log.Printf("Добавление страницы %v", sheet.Name)
	if err != nil {
		log.Printf("Ошибка добаления страницы %v", err)
	}

	// добавляем ячейки в строку (следуют одна за одной)
	for i := 1; i <= opts.N; i++ {
		if i == 1 {
			// создаем новую строку в странице
			row = sheet.AddRow()
			row.AddCell()
		}

		row.AddCell().SetInt(i)
	}

	// добавляем ячейки в столбец (следуют одна за одной)
	for i := 1; i <= opts.N; i++ {
		row = sheet.AddRow()
		row.AddCell().SetInt(i)
	}

	// сохранение
	err = file.Save(pwdDir + "/" + opts.FileSAVE)
	if err != nil {
		log.Printf("Ошибка сохранения файла %v", err)
	}
	log.Printf("Сохранение в %v", opts.FileSAVE)

	// готово
	log.Println("Готово")
}
