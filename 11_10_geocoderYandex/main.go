package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	// Широта и долгота
	resp, err := http.Get("https://geocode-maps.yandex.ru/1.x/?geocode=" + "Москва")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	point, err := regexp.Compile(`<pos>\d\d\.\d\d\d\d\d\d \d\d\.\d\d\d\d\d\d</pos>`)
	fmt.Println(string(point.Find(body)))
}
