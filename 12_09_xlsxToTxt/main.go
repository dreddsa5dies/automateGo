package main

import (
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
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

	// готово
	log.Println("Готово")
}
