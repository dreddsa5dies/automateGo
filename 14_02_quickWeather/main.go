// выводит прогноз погоды для заданного населенного пункта
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	golog    bool
	location string
)

func init() {
	flag.BoolVar(&golog, "log", false, "Use logfile")

	flag.StringVar(&location, "location", "London", "Location")
}

func main() {
	// разбор флагов
	flag.Parse()

	if golog {
		// создание файла log для записи ошибок
		fLog, err := os.OpenFile(`./.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
		if err != nil {
			log.Fatalln(err)
		}
		defer fLog.Close()

		// запись ошибок и инфы в файл и вывод
		log.SetOutput(io.MultiWriter(fLog, os.Stdout))
	}

	// GET API + key
	url := `http://api.openweathermap.org/data/2.5/weather?q=` + location + `&appid=109900a006b21159601c9b503d7c419c`
	log.Printf("url = %v", url)
	// test
	// url := `http://samples.openweathermap.org/data/2.5/forecast?q=London&appid=b1b15e88fa797225412429c1c50c122a1`

	// обращение к API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	log.Printf("status response = %v", resp.Status)

	// запись ответа в переменную
	body, err := ioutil.ReadAll(resp.Body)

	// создание переменной для хранения ответа
	// если структура ответа неизвестна
	var p map[string]interface{}

	// декодирование []byte в интерфейс
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	s := p["weather"].([]interface{})
	ss := s[0].(map[string]interface{})
	fmt.Println("Погода сегодня в " + location + ":")
	fmt.Printf("%v - %v\n", ss["main"], ss["description"])

	// для тестового работает
	/*
		s := p["list"].([]interface{})
		ss := s[0].(map[string]interface{})
		sss := ss["weather"].([]interface{})
		ssss := sss[0].(map[string]interface{})

		fmt.Println("Погода сегодня в " + location + ":")
		fmt.Printf("%v - %v\n", ssss["main"], ssss["description"])
	*/

}
