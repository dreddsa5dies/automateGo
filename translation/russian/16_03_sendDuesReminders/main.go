// считывание xlsx и отправка писем по SMTP
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {
	if len(os.Args) == 1 {
		println(`Запуск программы:	./sendDuesReminders PASSWD_EMAIL
		Рассылает сообщения на основании сведений из электронной таблицы об уплате взносов.`)
	} else {
		// создание файла log для записи ошибок
		fLog, err := os.OpenFile(`./.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
		if err != nil {
			log.Fatalln(err)
		}
		defer fLog.Close()

		// запись ошибок и инфы в файл и вывод
		log.SetOutput(io.MultiWriter(fLog, os.Stdout))

		excelFileName := "./duesRecords.xlsx"

		log.Printf("Открытие рабочей книги: %v", excelFileName)

		wb, err := xlsx.OpenFile(excelFileName)
		if err != nil {
			log.Fatalf("Ошибка открытия книги %v", err)
		}

		sheet := wb.Sheets[0]
		log.Printf("Открытие страницы: %v", sheet.Name)

		// сколько всего колонок
		// определение заполненных колонок
		var lastCol int
		for r := 0; r <= sheet.MaxCol; r++ {
			c := sheet.Cell(0, r).Value
			if c == "" {
				break
			}
			lastCol++
		}
		log.Printf("Всего колонок: %v", lastCol)

		// TODO: Проверить состояние уплаты взносов для каждого члена клуба
		// определение заполненных строк
		var maxRows int
		for r := 0; r <= sheet.MaxRow; r++ {
			c := sheet.Cell(r, 0).Value
			if c == "" {
				break
			}
			maxRows++
		}
		// заполнение карты
		unpaidMembers := make(map[string][]string)
		for r := 2; r < maxRows; r++ {
			payment := sheet.Cell(r, lastCol).Value
			if payment != "paid" {
				name := sheet.Cell(r, 0).Value
				email := sheet.Cell(r, 1).Value
				// определить месяц неуплаты
				var lastColName int
				for j := 0; j <= lastCol; j++ {
					c := sheet.Cell(r, j).Value
					if c == "" {
						break
					}
					lastColName++
				}
				latestMonth := sheet.Cell(0, lastColName-1)
				if err != nil {
					log.Fatalf("Ошибка получения данных %v", err)
				}
				log.Printf("Крайний месяц неуплаты: %v", latestMonth.Value)
				var allDataForName []string
				allDataForName = append(allDataForName, email)
				allDataForName = append(allDataForName, latestMonth.Value)
				unpaidMembers[name] = allDataForName
				log.Printf("Не заплатил: %v", name)
			}
		}

		// Войти в учетную запись email
		// Отправить сообщения с напоминанием об уплате
		for names, data := range unpaidMembers {
			log.Printf("Отправка для %v\t%v", names, data[0])
			// от кого кому и что
			from := mail.Address{Name: "gotestsmtp", Address: "gotestsmtp@gmail.com"}
			to := mail.Address{Name: names, Address: data[0]}

			// latestMonth в моем случае не значение
			body := "Уважаемый " + names + "\nОплатите " + data[1] + "\nС уважением, Виктор.\n"
			subject := "Вы забыли уплатить за " + data[1]

			// удаленный сервер Google (обязательно 587 порт) и данные аутотентификации
			smtpserver := "smtp.gmail.com:587"
			auth := smtp.PlainAuth("", "gotestsmtp@gmail.com", os.Args[1], "smtp.gmail.com")

			// установка заголовка письма
			header := make(map[string]string)
			header["From"] = from.String()
			header["To"] = to.String()
			header["Subject"] = subject

			// для него тело письма
			message := ""
			for k, v := range header {
				message += fmt.Sprintf("%s: %s\r\n", k, v)
			}
			message += "\r\n" + body

			// коннект с SMTP сервером
			c, err := smtp.Dial(smtpserver)
			if err != nil {
				log.Panic(err)
			}

			// без сертификата TLS Gmail не пропустит
			host, _, _ := net.SplitHostPort(smtpserver)
			tlc := &tls.Config{
				InsecureSkipVerify: true,
				ServerName:         host,
			}
			if err = c.StartTLS(tlc); err != nil {
				log.Panic(err)
			}

			// аутотентификация
			if err = c.Auth(auth); err != nil {
				log.Panic(err)
			}

			// отправка для КОГО и ЧТО
			if err = c.Mail(from.Address); err != nil {
				log.Panic(err)
			}
			if err = c.Rcpt(to.Address); err != nil {
				log.Panic(err)
			}

			// Само письмо
			w, err := c.Data()
			if err != nil {
				log.Panic(err)
			}
			_, err = w.Write([]byte(message))
			if err != nil {
				log.Panic(err)
			}
			err = w.Close()
			if err != nil {
				log.Panic(err)
			}
			c.Quit()
			log.Printf("Отправка для %v завершена", names)
		}

		log.Println("Готово")
	}
}
