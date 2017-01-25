package main

/*
   Программа переименовывает файлы, имена которых включают даты,
   указанные в американском формате (ММ-ДД-ГГГГ), приводя в соответствие с
   европейским форматом дат (ДД-ММ-ГГГГ)
*/

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) == 1 {
		// вывод справки
		reference := `
ИМЯ

renameDate -    переименовывает файлы, имена которых включают даты,
                указанные в американском формате (ММ-ДД-ГГГГ), приводя 
                в соответствие с европейским форматом дат (ДД-ММ-ГГГГ). 

СИНТАКСИС

renameDate КАТАЛОГ  

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
	} else {
		// TODO: создания паттерна рег выражения для файлов
		regStr, _ := regexp.Compile(`(.*)(\d{1,2})-(\d{1,2})-(19|20\d\d)(.*)`)
		// TODO: цикл по файлам каталога
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
		// сам цикл
		for _, fi := range fileInfos {
			strFile := fi.Name()
			// TODO: пропуск файлов не соответствующих регулярному выражению
			if regStr.MatchString(strFile) {
				// TODO: получения отдельных частей имен файла
				findRegexp := regStr.FindStringSubmatch(strFile)
				fmt.Println(findRegexp)
				for i := range findRegexp {
					fmt.Println(findRegexp[i] + ",")
				}
			}
		}
		// TODO: сформировать имена, соотв Европейскому формату
		// TODO: получить полные абсолютные пути к файлам
		// TODO: переименование файлов
	}
}

/*
Функции для перемещения и копирования файлов:
//перемещаем файл
func moveFile(src, dst string) error {
 err := os.Rename(src, dst)
 if err != nil {
  return err
 }
 return nil
}
//копируем файл
func copyFile(src string, dst string) (err error) {
 sourcefile, err := os.Open(src)
 if err != nil {
  return err
 }
 defer sourcefile.Close()
 destfile, err := os.Create(dst)
 if err != nil {
  return err
 }
 //копируем содержимое и проверяем коды ошибок
 _, err = io.Copy(destfile, sourcefile)
 if closeErr := destfile.Close(); err == nil {
  //если ошибки в io.Copy нет, то берем ошибку от destfile.Close(), если она была
  err = closeErr
 }
 if err != nil {
  return err
 }
 sourceinfo, err := os.Stat(src)
 if err == nil {
  err = os.Chmod(dst, sourceinfo.Mode())
 }
 return err
}
*/
