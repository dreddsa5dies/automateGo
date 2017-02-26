package main

/*
   Выборочное копирование: обход дерева каталогов с целью отбора файлов с заданным расширением,
   копирование в новую папку.
*/

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const reference string = `
ИМЯ

selectiveBackup - Копирует файлы из папок с заданным расширением. 

СИНТАКСИС

selectiveBackup [pdf/jpg] КАТАЛОГ

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
		// TODO: выбор расширения
		// парсинг флагов
		if len(os.Args) == 3 && os.Args[1] == "pdf" {
			// TODO: обход указанной директории
			readdir(os.Args[2])
		} else if os.Args[1] == "jpg" {
			readdir(os.Args[2])
		} else {
			fmt.Println(reference)
		}
		fmt.Println("Готово")
	}
}

func readdir(dir string) {
	// открываем директорию
	dh, err := os.Open(dir)
	if err != nil {
		return
	}
	defer dh.Close()
	// считывание списка файлов
	for {
		fis, err := dh.Readdir(10)
		if err == io.EOF {
			break
		}
		// вывод списка
		for _, fi := range fis {
			// тут доложно быть копирование
			if strings.HasSuffix(fi.Name(), os.Args[1]) {
				// TODO: копирование выбранных файлов
				fmt.Printf("Копирование %s/%s\n", dir, fi.Name())
				srcFile := dir + "/" + fi.Name()
				// адрес назначения можно подкорректировать
				dstFile := "/tmp/" + fi.Name()
				copyFile(srcFile, dstFile)
			}
			// рекурсивный проход по поддиректориям
			if fi.IsDir() {
				readdir(dir + "/" + fi.Name())
			}
		}
	}
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
