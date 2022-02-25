// Открывает первые 5 результатов с помощью Google
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/opesun/goquery"
	"github.com/toqueteos/webbrowser"

	"strings"
)

const reference string = `
ИМЯ

luckyGoogle - Открывает первые 5 результатов с помощью Google.

СИНТАКСИС

luckyGoogle ТЕКСТ

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
	if len(os.Args) == 1 {
		// вывод справки
		fmt.Println(reference)
	} else {
		pwdDir, _ := os.Getwd()
		// создание файла log
		// и нормальная обработка лога
		fLog, err := os.OpenFile(pwdDir+"/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
		check(err, fLog)

		// формирование запроса
		var search string
		for i := 1; i < len(os.Args); i++ {
			// добавление пробелов
			search += os.Args[i] + `%20`
		}
		// обрезать пробелы в строке
		strings.Trim(search, `%20`)

		// запрос по url
		resp, err := http.Get("https://www.google.ru/search?q=" + search)
		check(err, fLog)
		// отложенное закрытие коннекта
		defer resp.Body.Close()

		// парсинг ответа
		x, err := goquery.Parse(resp.Body)
		check(err, fLog)

		// храниение итоговых ссылок
		var urlsSearchs []string

		// формирование нормальной ссылки
		for _, i := range x.Find("h3").HtmlAll() {
			// обрезка html
			j := strings.TrimLeft(i, `<a href="/url?q=`)
			// надо убрать "левый" код в ссылке
			k := strings.Split(j, `&amp;sa=U&amp;ved=`)
			// итоговая ссылка готова
			urlsSearchs = append(urlsSearchs, "h"+k[0])
		}

		// Google luck
		// если ссылок меньше 5ти
		if len(urlsSearchs) < 5 {
			for i := 0; i < len(urlsSearchs); i++ {
				webbrowser.Open(urlsSearchs[i])
			}
		} else {
			// тут ограничение на 5ть ссылок
			for i := 0; i < 5; i++ {
				webbrowser.Open(urlsSearchs[i])
			}
		}
	}
}

// err check to log
func check(err error, fLog *os.File) {
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))
}
