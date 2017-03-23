// выводит прогноз погоды для заданного населенного пункта
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reflect"

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
	// долго копался и лень все описывать, поэтому встречайте
	p := make(map[string]interface{})

	// декодирование ответа API
	if err = json.NewDecoder(resp.Body).Decode(&p); err != nil {
		log.Fatalln(err)
	}
	w := p["list"]

	// заморочки с интерфейсами
	s := reflect.ValueOf(w)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	for i := 0; i < len(ret); i++ {
		fmt.Println(ret[i])
	}
}
