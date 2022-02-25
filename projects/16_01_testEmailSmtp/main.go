// smtp + ssl (для Google - 500 сообщений в день)
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

func main() {
	// от кого кому и что
	from := mail.Address{Name: "gotestsmtp", Address: "gotestsmtp@gmail.com"}
	to := mail.Address{Name: "gotest", Address: "gotest@lenta.ru"}

	body := "this is the body line1.\nthis is the body line2.\nthis is the body line3.\n"
	subject := "Тестовое Golang"

	// удаленный сервер Google (обязательно 587 порт) и данные аутотентификации
	smtpserver := "smtp.gmail.com:587"
	auth := smtp.PlainAuth("", "gotestsmtp@gmail.com", "PASSWORD", "smtp.gmail.com")

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
}
