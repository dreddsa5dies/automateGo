// выводит прогноз погоды для заданного населенного пункта
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Location string `short:"l" long:"location" default:"London" description:"Название населенного пункта"`
}

// разобрать всю структуру ответа JSON
type coord struct {
	Lat json.Number `json:"lat"`
	Lon json.Number `json:"lon"`
}

type city struct {
	ID    json.Number `json:"id"`
	Name  string      `json:"name"`
	Coord coord       `json:"coord"`
}

// пока закончил тут
type weather struct {
	ID   json.Number `json:"id"`
	Main string      `json:"main"`
}

type mainData struct {
	Temp     json.Number `json:"temp"`
	TempMin  json.Number `json:"temp_min"`
	TempMax  json.Number `json:"temp_max"`
	Pressure json.Number `json:"pressure"`
	SeaLevel json.Number `json:"sea_level"`
	Humidity json.Number `json:"humidity"`
	TempKf   json.Number `json:"temp_kf"`
	Weather  weather     `json:"weather"`
}

type data struct {
	Dt   json.Number `json:"dt"`
	Main mainData    `json:"main"`
}

type weatherData struct {
	Cod     string      `json:"cod"`
	Message json.Number `json:"message"`
	Cnt     json.Number `json:"cnt"`
	List    []data      `json:"list"`
	City    city        `json:"city"`
	Country string      `json:"country"`
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

	// формат вывода log
	logger := log.New(fLog, "testJSON ", log.LstdFlags|log.Lshortfile)

	// GET API + key
	url := `http://api.openweathermap.org/data/2.5/forecast/city?q=` + opts.Location + `&APPID=109900a006b21159601c9b503d7c419c`

	// обращение к API
	resp, err := http.Get(url)
	if err != nil {
		logger.Fatalln(err)
	}
	defer resp.Body.Close()

	// создание переменной для хранения ответа
	// необходимо составлять структуру для хранения ответа
	var p weatherData

	// декодирование ответа API
	if err = json.NewDecoder(resp.Body).Decode(&p); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(p)
}
