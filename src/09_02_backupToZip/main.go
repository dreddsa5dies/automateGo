package main

/*
   Копирует папку вместе со всем содержимым
   в zip-файл с инкреминтируемым номером копии в имени файла
*/

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		// вывод справки
		reference := `
ИМЯ

backupToZip -   Копирует папку вместе со всем содержимым в zip-файл 
		с инкреминтируемым номером копии в имени файла. 

СИНТАКСИС

backupToZip КАТАЛОГ  

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
		// TODO: проверка на абсолютный путь к файлу
		filePath, err := filepath.Abs(os.Args[1])
		if err != nil {
			return
		}
		fmt.Println(filePath)
		var zipFileName string
		// TODO: определение имени которое следует присвоить файлу (исходя из существующих)
		number := 1
		for true {
			zipFileName = filePath + "_" + strconv.Itoa(number) + ".zip"
			// не работает
			if _, err := os.Stat(zipFileName); os.IsNotExist(err) {
				break
			}
			number++
		}
		fmt.Println("! " + zipFileName)
		// TODO: создание zip-файла
		// TODO: обойти все дерево папки и сжать файлы в каждой папке
		fmt.Println("Готово")
	}
}

/*
// архивирование
// zipit("/tmp/documents", "/tmp/backup.zip")
// zipit("/tmp/report.txt", "/tmp/report-2015.zip")
func zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
*/
