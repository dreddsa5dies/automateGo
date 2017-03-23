// выводит прогноз погоды для заданного населенного пункта
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Location string `short:"l" long:"location" default:"London" description:"Название населенного пункта"`
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

	url := `http://api.openweathermap.org/data/2.5/forecast/city?q=` + opts.Location + `&APPID=109900a006b21159601c9b503d7c419c`

	resp, err := http.Get(url)
	if err != nil {
		logger.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		logger.Fatalln(err)
	}

	fmt.Println(f)
}
