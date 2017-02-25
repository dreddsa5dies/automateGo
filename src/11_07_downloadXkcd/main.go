package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"strings"

	"github.com/opesun/goquery"
)

const reference string = `
ИМЯ

downloadXkcd - загружает комиксы XKCD.com

СИНТАКСИС

downloadXkcd

АВТОР

Виктор Соловьев 

СООБЩЕНИЕ ОБ ОШИБКАХ

Об ошибках сообщайте по адресу <viktor.vladimirovich.solovev@gmail.com>.  

АВТОРСКИЕ ПРАВА

Copyright 2017 Viktor Solovev

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.`

func main() {
	if len(os.Args) > 1 {
		// вывод справки
		fmt.Println(reference)
	} else {
		pwdDir, _ := os.Getwd()
		// создание файла log
		// и нормальная обработка лога
		fLog, err := os.OpenFile(pwdDir+"/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
		check(err, fLog)

		// исходный url
		url := `http://xkcd.com`

		err = os.Mkdir(pwdDir+"/xkcd", 0775)
		check(err, fLog)

		// for i := 0; i < 5; i++ {
		// TODO: Загрузить страницу
		// запрос по url
		resp, err := http.Get(url)
		check(err, fLog)
		// отложенное закрытие коннекта
		defer resp.Body.Close()

		// TODO: Найти URL комикса
		// парсинг ответа
		x, err := goquery.Parse(resp.Body)
		check(err, fLog)

		// нахождение ссылки
		regStr, _ := regexp.Compile(`comics`)
		for _, i := range x.Find("img").Attrs("src") {
			if regStr.MatchString(i) {
				// формирование имени файла
				nameFile := strings.Split(i, "/comics/")
				// создание файла для загрузки картинки
				name := pwdDir + "/xkcd/" + nameFile[1]
				fSave, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0640)
				check(err, fLog)

				// TODO: Загрузить комикс
				respImg, err := http.Get("http:" + i)
				check(err, fLog)
				// отложенное закрытие коннекта
				defer respImg.Body.Close()

				// запись ответа в переменную
				bodyImg, err := ioutil.ReadAll(respImg.Body)
				check(err, fLog)

				// TODO: Сохранить изображение
				fSave.Write(bodyImg)
			}
		}

		// TODO: Получить URL для prev
		//}
		log.Println("Готово")
		log.SetOutput(io.MultiWriter(fLog, os.Stdout))
	}
}

// err check to log
func check(err error, fLog *os.File) {
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))
}
