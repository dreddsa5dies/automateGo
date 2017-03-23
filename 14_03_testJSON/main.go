//считывание JSON
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// в какой папке исполняемы файл
	pwdDir, _ := os.Getwd()

	// создание файла log для записи ошибок
	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	// запись ошибок и инфы в файл
	log.SetOutput(fLog)

	// структура данных
	type Key struct {
		State, Country string
	}
	type SaveData struct {
		Name   Key     `json:"state and country name"`
		Pop    float64 `json:"populations"`
		Tracts int     `json:"tracts"`
	}

	log.Println("Начало")

	data, err := ioutil.ReadFile("exitData.json")
	if err != nil {
		log.Fatalln("Ошибка открытия файла")
	}
	log.Println(data)

	var jsontype SaveData
	// только 1 строчку
	err = json.Unmarshal(data, &jsontype)
	if err != nil {
		log.Fatalln("Ошибка декодирования")
	}

	log.Println(jsontype.Name)
	log.Println(jsontype.Pop)
	log.Println(jsontype.Tracts)
}
