// регулярное вырадение проверки пароля
package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Использование: %s password\n", os.Args[0])
		os.Exit(1)
	}

	s := os.Args[1]

	regStr, _ := regexp.Compile(`([0-9a-zA-Z]){8,}`)

	if regStr.MatchString(s) {
		fmt.Println("Pass ok")
	} else {
		fmt.Println("Bad pass")
	}

	os.Exit(0)
}
