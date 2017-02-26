package main

/*
поиск в *.txt по регулярному выражению
*/

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		// вывод справки
		reference := `
ИМЯ

regexpTxt -    поиск во всех *.txt в папке по регулярному выражению. 

СИНТАКСИС

regexpTxt КАТАЛОГ  

АВТОР

Виктор Соловьев 

СООБЩЕНИЕ ОБ ОШИБКАХ

Об ошибках сообщайте по адресу <viktor.vladimirovich.solovev@gmail.com>.  

АВТОРСКИЕ ПРАВА

Copyright 2017 Viktor Solovev

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.`
		fmt.Println(reference)
	} else if len(os.Args) == 2 {
		//открытие папки
		dir, err := os.Open(os.Args[1])
		if err != nil {
			return
		}
		defer dir.Close()

		// список файлов
		fileInfos, err := dir.Readdir(-1)
		if err != nil {
			return
		}
		// создание шаблона регулярного выражения
		fmt.Print("Регулярное выражение> ")
		regexTxt, _ := regexp.Compile(scan())

		for _, fi := range fileInfos {
			// проверка на .txt
			if strings.HasSuffix(fi.Name(), ".txt") {
				// проверка корректности имени файла и его открытие
				file, err := ioutil.ReadFile(fi.Name())
				if err != nil {
					fmt.Println(err)
				}
				// формируем строку из []byte
				strFile := string(file)
				// поиск по регулярному выражению
				var founds []string
				if regexTxt.MatchString(strFile) {
					founds = regexTxt.FindAllString(strFile, -1)
				}

				// вывод данных
				if len(founds) > 0 {
					fmt.Printf("Founds in %v:\n", fi.Name())
					fmt.Println(strings.Join(founds, "\n"))
				} else {
					fmt.Println("Not found.")
				}
			}
		}
	}
}

// функция ввода данных
func scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
