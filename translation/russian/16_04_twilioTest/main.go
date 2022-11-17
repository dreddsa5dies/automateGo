// SMS cannot be sent to landline destination number. The Twilio REST API will throw a 400 response with error code 21614, the message will not appear in the logs and the account will not be charged.
package main

import (
	"io"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/sfreiberg/gotwilio"
)

var opts struct {
	SID  string `short:"s" long:"sid" default:"" description:"accountSid"`
	AUTH string `short:"a" long:"auth" default:"" description:"authToken"`
}

func main() {
	if len(os.Args) == 1 {
		println(`Please ./twilioTest -h`)
	} else {
		// разбор флагов
		flags.Parse(&opts)

		// создание файла log для записи ошибок
		fLog, err := os.OpenFile(`./.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
		if err != nil {
			log.Fatalln(err)
		}
		defer fLog.Close()

		// запись ошибок и инфы в файл и вывод
		log.SetOutput(io.MultiWriter(fLog, os.Stdout))

		twilio := gotwilio.NewTwilioClient(opts.SID, opts.AUTH)

		from := "+12013807451"
		to := "+79818500301"
		message := "Welcome to gotwilio!"
		twilio.SendSMS(from, to, message, "", "")

		log.Println("Готово")
	}
}
