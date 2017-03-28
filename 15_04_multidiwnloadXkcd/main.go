//загружает комиксы XKCD.com многопоточно, но кривовато
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"

	"strings"

	"strconv"

	"sort"

	"github.com/opesun/goquery"
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

	// исходный url
	url := `http://xkcd.com`

	err = os.Mkdir(pwdDir+"/xkcd", 0775)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO: Загрузить страницу
	// запрос по url
	resp, err := http.Get(url)
	log.Printf("Загружается страница:	%v", url)
	if err != nil {
		log.Fatalln(err)
	}
	// отложенное закрытие коннекта
	defer resp.Body.Close()

	// TODO: Найти URL комикса
	// парсинг ответа
	x, err := goquery.Parse(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO: Получить URL для prev
	// ищу ссылочки на старые комиксы
	var urlInt []int
	regLastLink, _ := regexp.Compile(`\/[0-9]{1,}\/`)
	for _, i := range x.Find("a").Attrs("href") {
		if regLastLink.MatchString(i) {
			a, err := strconv.Atoi(strings.Trim(i, "/"))
			if err != nil {
				log.Fatalln(err)
			}
			urlInt = append(urlInt, a)
		}
	}
	sort.Ints(urlInt)
	// предыдущая ссылка
	log.Println(urlInt[len(urlInt)-1])

	// формирование массива ссылок
	var massUrls []string
	for k := urlInt[len(urlInt)-1] + 1; k >= 1; k-- {
		massUrls = append(massUrls, url+"/"+strconv.Itoa(k))
	}

	// TODO: многопоточность
	byteResponse := make(chan []byte)
	var wg sync.WaitGroup
	// обязательно ожидание всех горутин
	wg.Add(len(massUrls))
	for _, urls := range massUrls {
		go func(urls string) {
			defer wg.Done()
			resp, err := http.Get(urls)
			log.Printf("Загружается страница:	%v", urls)
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()

			// TODO: Найти URL комикса
			// парсинг ответа
			x, err := goquery.Parse(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			// нахождение ссылки
			regStr, _ := regexp.Compile(`comics`)
			for _, i := range x.Find("img").Attrs("src") {
				if regStr.MatchString(i) {
					// TODO: Загрузить комикс
					respImg, err := http.Get("http:" + i)
					log.Printf("Загружается изображение:	%v", "http:"+i)
					if err != nil {
						log.Fatalln(err)
					}
					// отложенное закрытие коннекта
					defer respImg.Body.Close()

					// запись ответа в переменную
					bodyImg, err := ioutil.ReadAll(respImg.Body)
					if err != nil {
						log.Fatalln(err)
					}

					byteResponse <- bodyImg
				}
			}

			time.Sleep(100 * time.Millisecond)
		}(urls)
	}

	go func() {
		for response := range byteResponse {
			// создание файла для загрузки картинки
			name := pwdDir + "/xkcd/" + time.Now().String() + ".png"
			fSave, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0640)
			if err != nil {
				log.Fatalln(err)
			}
			defer fSave.Close()
			log.Printf("Сохрание в:	%v", name)
			fSave.Write(response)
		}
	}()
	wg.Wait()

	log.Println("Готово")
}
