package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

var code = `
package main
import "fmt"

func main() {
	fmt.Println("Hello WebDriver!\n")
}
`

func main() {
	pwdDir, _ := os.Getwd()
	// создание файла log
	// и нормальная обработка лога
	fLog, err := os.OpenFile(pwdDir+"/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	check(err, fLog)

	// FireFox driver without specific version
	// *** Add gecko driver here if necessary (see notes above.) ***
	caps := selenium.Capabilities{
		"browserName":            "firefox",
		"webdriver.gecko.driver": "$GOPATH/bin/geckodriver",
	}
	wd, err := selenium.NewRemote(caps, "")
	check(err, fLog)
	defer wd.Quit()

	// Get simple playground interface
	wd.Get("http://play.golang.org/?simple=1")

	// Enter code in textarea
	elem, _ := wd.FindElement(selenium.ByCSSSelector, "#code")
	elem.Clear()
	elem.SendKeys(code)

	// Click the run button
	btn, _ := wd.FindElement(selenium.ByCSSSelector, "#run")
	btn.Click()

	// Get the result
	div, _ := wd.FindElement(selenium.ByCSSSelector, "#output")

	output := ""
	// Wait for run to finish
	for {
		output, _ = div.Text()
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("Got: %s\n", output)
}

// err check to log
func check(err error, fLog *os.File) {
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))
}
