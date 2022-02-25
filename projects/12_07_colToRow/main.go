// замена колонок на с троки XLSX
package main

import (
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/tealeg/xlsx"
)

var opts struct {
	FileSAVE string `short:"s" long:"save" default:"save_table.xlsx" description:"Сохраненный файл таблицы"`
	FileOPEN string `short:"o" long:"open" default:"table.xlsx" description:"Файл таблицы"`
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

	// открываем файл
	excelFileName := pwdDir + "/" + opts.FileOPEN
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Printf("Ошибка открытия файла %v", err)
	}
	// Имя страницы открытого файла
	i := xlFile.Sheets[0]
	log.Printf("Открытие страницы %v файла %v", i.Name, opts.FileOPEN)

	// создаем новый файл
	file := xlsx.NewFile()
	// добавляем страницу
	sheet, err := file.AddSheet("Sheet")
	log.Printf("Добавление страницы %v", sheet.Name)
	if err != nil {
		log.Printf("Ошибка добаления страницы %v", err)
	}

	// блок чтения и записи
	for l := 0; l < i.MaxCol; l++ {
		// сколько всего строк
		k := i.MaxRow
		// до добавить строку (т.е. перевести столбы в строки)
		sheetRow := sheet.AddRow()
		// до максимума строк пройти
		for u := 0; u < k; u++ {
			// по столбцу l
			sd := i.Cell(u, l)
			// и вывести как строку
			ss, _ := sd.String()
			// добавить в ячейку нового файла
			sheetRow.AddCell().SetValue(ss)
		}
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
