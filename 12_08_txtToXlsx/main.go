package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/tealeg/xlsx"
)

var opts struct {
	FileSAVE string `short:"s" long:"save" default:"saveTable.xlsx" description:"Сохраненный файл таблицы"`
}

func main() {
	// разбор флагов
	flags.Parse(&opts)

	// в какой папке исполняемый файл
	pwdDir, _ := os.Getwd()

	// создание файла log для записи ошибок
	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	// запись ошибок и инфы в файл
	log.SetOutput(fLog)

	// создаем новый файл
	file := xlsx.NewFile()
	// добавляем страницу
	sheet, err := file.AddSheet("Sheet")
	log.Printf("Добавление страницы %v", sheet.Name)
	if err != nil {
		log.Printf("Ошибка добаления страницы %v", err)
	}

	// TODO: считывание всех txt файлов в каталоге
	//открытие папки
	dir, err := os.Open(pwdDir)
	if err != nil {
		return
	}
	defer dir.Close()

	// список файлов
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	// хеш для хранения строк из файлов
	dataFiles := make(map[string][]string)

	// для всех файлов в папке
	for _, fi := range fileInfos {
		// проверка на .txt
		if strings.HasSuffix(fi.Name(), ".txt") {
			log.Printf("Считывание %v", fi.Name())
			// построчное считывание
			fileOpen, err := os.Open(fi.Name())
			if err != nil {
				log.Fatalln(err)
			}
			// массив для хранения строк
			var arr []string
			scanner := bufio.NewScanner(fileOpen)
			for scanner.Scan() {
				// строки в массив
				arr = append(arr, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			defer fileOpen.Close()

			// массив в hash
			dataFiles[fi.Name()] = arr
		}
	}

	// TODO: запись всего считанного в файлах
	// счетчик для столбцов
	i := 0
	for keys, prs := range dataFiles {
		// запись наименования файла в заголовок
		sheet.Cell(0, i).SetValue(keys)
		// запись в строку
		// счетчик для строк
		l := 1
		for _, str := range prs {
			// запись данных в строки столбцов
			sheet.Cell(l, i).SetValue(str)
			l++
		}
		i++
	}

	// сохранение
	err = file.Save(pwdDir + "/" + opts.FileSAVE)
	if err != nil {
		log.Printf("Ошибка сохранения файла %v", err)
	}
	log.Printf("Сохранение в %v", opts.FileSAVE)

	// готово
	log.Println("Готово")
}
