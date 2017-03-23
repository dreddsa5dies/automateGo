//Selenium click
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"time"

	"github.com/tebeka/selenium"
)

func main() {
	pwdDir, _ := os.Getwd()
	fLog, err := os.OpenFile(pwdDir+"/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	check(err, fLog)

	// настройки для Selenium
	// обязательно сервер Selenium запускать
	// java -jar /home/dreddsa/go/src/github.com/dreddsa5dies/automateGo/src/11_08_seleniumGo/selenium-server-standalone-3.1.0.jar
	caps := selenium.Capabilities{
		"browserName":            "firefox",
		"webdriver.gecko.driver": "$GOPATH/bin/geckodriver",
	}
	wd, err := selenium.NewRemote(caps, "")
	check(err, fLog)
	defer wd.Quit()

	// get url
	wd.Get("http://inventwithpython.com")

	// поиск элементов с классом bookcover
	elem, err := wd.FindElement(selenium.ByClassName, "bookcover")
	check(err, fLog)

	// имя тега <img>
	a, err := elem.TagName()
	check(err, fLog)
	fmt.Printf("Найден элемент %v с данным имененм класса!\n", a)

	linkElem, err := wd.FindElement(selenium.ByLinkText, "online course of this book on Udemy")
	check(err, fLog)

	linkElem.Click()
	// подождать, чтобы прогрузилась ссылка
	time.Sleep(time.Millisecond * 10000)
}

// err check to log
func check(err error, fLog *os.File) {
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))
}
