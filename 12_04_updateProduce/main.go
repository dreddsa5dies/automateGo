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

	// готово
	log.Println("Готово")
}
