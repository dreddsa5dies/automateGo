package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	var s, latitude, longitude string
	s = scan()
	// Широта и долгота
	resp, err := http.Get("https://geocode-maps.yandex.ru/1.x/?geocode=" + s)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	point, err := regexp.Compile(`<pos>\d\d\.\d\d\d\d\d\d \d\d\.\d\d\d\d\d\d</pos>`)
	strFind := string(point.Find(body))
	strFind = strings.TrimLeft(strFind, "<pos>")
	strFind = strings.TrimRight(strFind, "</pos>")
	strDATA := strings.Split(strFind, " ")
	if len(strDATA) > 1 {
		latitude = strDATA[0]
		longitude = strDATA[1]
	} else {
		latitude, longitude = "неизвестно", "неизвестно"
	}
	fmt.Printf("Широта\t> %v\n", latitude)
	fmt.Printf("Долгота\t> %v\n", longitude)
}

func scan() string {
	fmt.Print("Ввод наименования объекта > ")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
