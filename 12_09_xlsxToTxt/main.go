package main

import (
	"log"
	"os"
	"strconv"

	flags "github.com/jessevdk/go-flags"
	"github.com/tealeg/xlsx"
)

var opts struct {
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

	excelFileName := pwdDir + "/" + opts.FileOPEN

	log.Printf("Открытие рабочей книги: %v", excelFileName)

	wb, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("Ошибка открытия книги %v", err)
	}

	// Имя страницы
	i := wb.Sheets[0]
	log.Printf("Открытие страницы %v", i.Name)

	// генерация файлов
	for m := 0; m < i.MaxCol; m++ {
		// создание
		strName := "saveFile" + strconv.Itoa(m+1) + ".txt"
		saveFile, err := os.Create(strName)
		log.Printf("Создание файла отчета %v", saveFile.Name())
		if err != nil {
			log.Fatalf("Ошибка создания файла %v", err)
		}
		defer saveFile.Close()

		// до максимума строк пройти
		for u := 0; u < i.MaxRow; u++ {
			// по столбцу
			sd := i.Cell(u, m)
			// и вывести как строку
			ss, _ := sd.String()
			// Запись
			saveFile.WriteString(ss)
			saveFile.WriteString("\n")
		}
	}

	// готово
	log.Println("Готово")
}
