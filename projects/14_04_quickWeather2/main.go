// разбор JSON с неизвестной структурой
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// создание файла log для записи ошибок
	fLog, err := os.OpenFile(`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	// запись ошибок и инфы в файл и вывод
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))

	file, err := ioutil.ReadFile("./resp.json")
	if err != nil {
		log.Fatalln("Ошибка открытия файла")
	}
	log.Println("Открытие JSON файла: resp.json")

	// создание переменной для хранения ответа
	// если структура ответа неизвестна
	var p interface{}

	// декодирование []byte в интерфейс
	err = json.Unmarshal(file, &p)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// функция вывода данных
	printJSON(p)

}

func printJSON(v interface{}) {
	// доступ по типу
	switch vv := v.(type) {
	// строки
	case string:
		fmt.Println("is string", vv)
	// числа
	case float64:
		fmt.Println("is float64", vv)
	// вложенные данные (рекурсия)
	case []interface{}:
		fmt.Println("is array:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	// неизвестные данные
	default:
		fmt.Println("Unknown type")
	}
}
