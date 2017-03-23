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

// Coordinates struct holds longitude and latitude data in returned
// JSON or as parameter data for requests using longitude and latitude.
type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// Sys struct contains general information about the request
// and the surrounding area for where the request was made.
type Sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

// Wind struct contains the speed and degree of the wind.
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

// Weather struct holds high-level, basic info on the returned
// data.
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main struct contains the temperates, humidity, pressure for the request.
type Main struct {
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
}

// Clouds struct holds data regarding cloud cover.
type Clouds struct {
	All int `json:"all"`
}

// CurrentWeatherData above for JSON to be unmarshaled into.
type CurrentWeatherData struct {
	GeoPos  Coordinates        `json:"coord"`
	Sys     Sys                `json:"sys"`
	Base    string             `json:"base"`
	Weather []Weather          `json:"weather"`
	Main    Main               `json:"main"`
	Wind    Wind               `json:"wind"`
	Clouds  Clouds             `json:"clouds"`
	Rain    map[string]float64 `json:"rain"`
	Snow    map[string]float64 `json:"snow"`
	Dt      int                `json:"dt"`
	ID      int                `json:"id"`
	Name    string             `json:"name"`
	Cod     string             `json:"cod"`
	Unit    string
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
	var p CurrentWeatherData

	// декодирование ответа API
	if err = json.NewDecoder(resp.Body).Decode(&p); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(p)
}
