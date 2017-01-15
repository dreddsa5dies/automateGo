package main

/*
замена ключевых слов
ADJECTIVE - прилагательное
NOUN - существительное
VERB - глагол
NOUN - наречие
*/

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		// вывод справки
		fmt.Println(`Замена ключевых слов в файле
	ADJECTIVE - прилагательное
NOUN - существительное
VERB - глагол
NOUN - наречие
Использование:
	madLib <name.txt>	адрес файла`)
	} else {
		if len(os.Args) == 3 {
			
}

// проверка на ошибки
func check(e error) {
	if e != nil {
		panic(e)
	}
}
