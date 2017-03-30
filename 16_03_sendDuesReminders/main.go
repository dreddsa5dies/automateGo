// считывание xlsx и отправка писем по SMTP
package main

import (
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {
	if len(os.Args) > 1 {
		println(`Start programm:	./sendDuesReminders`)
	} else {
		// создание файла log для записи ошибок
		fLog, err := os.OpenFile(`./.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
		if err != nil {
			log.Fatalln(err)
		}
		defer fLog.Close()

		// запись ошибок и инфы в файл
		log.SetOutput(fLog)

		excelFileName := "./duesRecords.xlsx"

		log.Printf("Открытие рабочей книги: %v", excelFileName)

		wb, err := xlsx.OpenFile(excelFileName)
		if err != nil {
			log.Fatalf("Ошибка открытия книги %v", err)
		}

		sheet := wb.Sheets[0]
		log.Printf("Открытие страницы: %v", sheet.Name)

		// сколько всего колонок
		k := sheet.MaxCol
		log.Printf("Всего колонок: %v", k)

		log.Println("Готово")
	}
}
