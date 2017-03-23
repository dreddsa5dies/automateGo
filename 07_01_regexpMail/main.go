// поиск по регулярному выражению email
package main

import (
	"fmt"
	"regexp"

	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	// создание шаблонов регулярных выражений
	// email regex
	regMail, _ := regexp.Compile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,6}`)

	// чтение из буфера
	// поиск по шаблонам и добавление в срез
	text, _ := clipboard.ReadAll()
	var mailAddr []string

	if regMail.MatchString(text) {
		mailAddr = regMail.FindAllString(text, -1)
	}

	// запись среза в буфер обмена
	if len(mailAddr) > 0 {
		clipboard.WriteAll(strings.Join(mailAddr, "\n"))
		fmt.Println("Copied to clipboard:")
		fmt.Println(strings.Join(mailAddr, "\n"))
	} else {
		fmt.Println("No email addresses found.")
	}
}
