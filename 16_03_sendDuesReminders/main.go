// считывание xlsx и отправка писем по SMTP
package main

import (
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {
	if len(os.Args) > 1 {
		println(`Запсук программы:	./sendDuesReminders
		Рассылает сообщения на основании сведений из электронной таблицы об уплате взносов.`)
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
		lastCol := sheet.MaxCol
		log.Printf("Всего колонок: %v", lastCol)
		latestMonth := sheet.Cell(0, lastCol)
		log.Printf("Крайний месяц: %v", latestMonth.Value)

		// TODO: Проверить состояние уплаты взносов для каждого члена клуба
		unpaidMembers := make(map[string]string)
		for r := 2; r <= sheet.MaxRow; r++ {
			payment := sheet.Cell(r, lastCol).Value
			if payment != "paid" {
				name := sheet.Cell(r, 0).Value
				email := sheet.Cell(r, 1).Value
				unpaidMembers[name] = email
				log.Printf("Не заплатил: %v", name)
			}
		}

		// TODO: Войти в учетную запись email

		// TODO: Отправить сообщения с напоминанием об уплате

		log.Println("Готово")
	}
}
