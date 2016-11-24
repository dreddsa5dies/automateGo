package main

import (
	"fmt"
	"regexp"
)

func main() {
	var (
		s string
	)

	regStr, _ := regexp.Compile(`([0-9a-zA-Z]){8,}`)

	fmt.Print("Pass: ")
	fmt.Scanf("%s", &s)

	if regStr.MatchString(s) {
		fmt.Println("Pass ok")
	} else {
		fmt.Println("Bad pass")
	}
}
