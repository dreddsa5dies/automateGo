// удаляет заголовки из всех CSV-файлов в текущем каталоге
package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// в какой папке исполняемый файл
	pwdDir, _ := os.Getwd()

	// создание файла log для записи ошибок
	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	// формат вывода log
	logger := log.New(fLog, "testCsv ", log.LstdFlags|log.Lshortfile)

	//открытие папки
	dir, err := os.Open(pwdDir)
	if err != nil {
		logger.Fatalln(err)
	}
	defer dir.Close()

	// список файлов
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		logger.Fatalln(err)
	}

	for _, fi := range fileInfos {
		// проверка на .csv
		if strings.HasSuffix(fi.Name(), ".csv") {
			logger.Printf("Считывание файла %v", fi.Name())
			// read file
			dat, err := ioutil.ReadFile(fi.Name())
			if err != nil {
				logger.Fatalf("Ошибка чтения файла: %v", err)
			}

			in := string(dat)

			// encoding csv
			r := csv.NewReader(strings.NewReader(in))

			// получение массива строк и ячеек
			records, err := r.ReadAll()
			if err != nil {
				logger.Fatalf("Ошибка чтения данных файла: %v", err)
			}

			// удаление 1й строки
			records = records[1:]

			// создание файла сохранения
			str := "new_" + fi.Name()
			saveFile, err := os.Create(str)
			if err != nil {
				logger.Fatalf("Ошибка чтения файла: %v", err)
			}
			defer saveFile.Close()
			logger.Printf("Создание файла сохранения %v", saveFile.Name())

			// запись данных в файл
			writer := csv.NewWriter(saveFile)
			writer.WriteAll(records)
			writer.Flush()
			logger.Println("Запись произведена")
		}
	}
	logger.Println("Готово")
}
