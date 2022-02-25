// вставка строк XLSX
package main

import (
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/tealeg/xlsx"
)

var opts struct {
	N        int    `short:"n" long:"number" default:"2" description:"Число N строк для вставки"`
	M        int    `short:"m" long:"row" default:"2" description:"Число M строки после которой вставляются строки"`
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
	rowsI := i.Rows

	// создаем новый файл
	file := xlsx.NewFile()
	// добавляем страницу
	sheet, err := file.AddSheet("Sheet")
	log.Printf("Добавление страницы %v", sheet.Name)
	if err != nil {
		log.Printf("Ошибка добаления страницы %v", err)
	}

	// блок записи новых строк
	for rowNum := 0; rowNum < i.MaxRow; rowNum++ {
		if rowNum != opts.M {
			// создаем овую строку в странице
			row := sheet.AddRow()
			cells := rowsI[rowNum].Cells
			for a := 0; a < len(cells); a++ {
				row.AddCell().SetValue(cells[a].Value)
			}
		} else {
			log.Printf("Добавление %v строк после строки №%v", opts.N, rowNum)
			for j := 0; j < opts.N; j++ {
				sheet.AddRow().AddCell()
			}
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
