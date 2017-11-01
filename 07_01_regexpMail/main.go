// поиск по регулярному выражению email
package main

import (
	"fmt"
	"os"
	"regexp"

	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	// helpline
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		fmt.Printf("Usage: %s\n", os.Args[0])
	} else {
		// create email regexp
		regMail, _ := regexp.Compile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,6}`)

		// read os buffer
		// find email regexp
		text, _ := clipboard.ReadAll()
		var mailAddr []string

		if regMail.MatchString(text) {
			mailAddr = regMail.FindAllString(text, -1)
		}

		// write on os buffer
		if len(mailAddr) > 0 {
			clipboard.WriteAll(strings.Join(mailAddr, "\n"))
			fmt.Println("Copied to clipboard:")
			fmt.Println(strings.Join(mailAddr, "\n"))
		} else {
			fmt.Println("No email addresses found.")
		}
	}
}
