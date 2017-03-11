package main

import (
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/tealeg/xlsx"
)

var opts struct {
	FileExel string `short:"o" long:"open" default:"censuspopdata.xlsx" description:"Региональная численность населения"`
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

	excelFileName := pwdDir + "/" + opts.FileExel

	log.Printf("Открытие рабочей книги: %v", excelFileName)

	wb, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("Ошибка открытия книги %v", err)
	}

	sheet := wb.Sheet["Population by Census Tract"]
	log.Printf("Открытие страницы: %v", sheet.Name)
}
