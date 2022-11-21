//Поиск исключительно больших файлов (> 100 MB).
package main

import (
	"fmt"
	"io"
	"os"
)

const reference string = `
ИМЯ

indLargeFiles - Поиск исключительно больших файлов (> 100 MB). 

СИНТАКСИС

indLargeFiles КАТАЛОГ

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
		// TODO: обход указанной директории
		readdir(os.Args[1])
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
			if fi.Size() > 104857600 {
				fmt.Printf("%s/%s\t%d MB\n", dir, fi.Name(), fi.Size()/1048576)
			}
			// рекурсивный проход по поддиректориям
			if fi.IsDir() {
				readdir(dir + "/" + fi.Name())
			}
		}
	}
}
