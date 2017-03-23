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

	// запись ошибок и инфы в файл
	log.SetOutput(fLog)

	//открытие папки
	dir, err := os.Open(pwdDir)
	if err != nil {
		log.Fatalln(err)
	}
	defer dir.Close()

	// список файлов
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		log.Fatalln(err)
	}

	for _, fi := range fileInfos {
		// проверка на .csv
		if strings.HasSuffix(fi.Name(), ".csv") {
			// read file
			dat, err := ioutil.ReadFile(fi.Name())
			if err != nil {
				log.Fatalf("Ошибка чтения файла: %v", err)
			}

			in := string(dat)

			// encoding csv
			r := csv.NewReader(strings.NewReader(in))

			// получение массива строк и ячеек
			records, err := r.ReadAll()
			if err != nil {
				log.Fatalf("Ошибка чтения данных файла: %v", err)
			}

			// удаление 1й строки
			records = records[1:]

			// создание файла сохранения
			str := "new" + fi.Name()
			saveFile, err := os.Create(str)
			if err != nil {
				log.Fatalf("Ошибка чтения файла: %v", err)
			}
			defer saveFile.Close()

			// запись данных в файл
			writer := csv.NewWriter(saveFile)
			writer.WriteAll(records)
			writer.Flush()
		}
	}
}
