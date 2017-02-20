package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const reference string = `
ИМЯ

googleSearchTerm - запрос поиска Google

СИНТАКСИС

googleSearchTerm ТЕКСТ

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

		var url string
		for i := 1; i < len(os.Args); i++ {
			url += os.Args[i] + "%20"
		}

		// запрос по url
		resp, err := http.Get("https://www.google.ru/webhp#newwindow=1&safe=off&q=" + url + "insite:sbis.ru")
		check(err, fLog)
		// отложенное закрытие коннекта
		defer resp.Body.Close()

		// забись ответа в переменную
		body, err := ioutil.ReadAll(resp.Body)
		check(err, fLog)

		// вывод содержимого
		fmt.Println(string(body))
	}
}

// err check to log
func check(err error, fLog *os.File) {
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))
}
