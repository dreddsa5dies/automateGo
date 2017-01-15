package main

/*
замена ключевых слов
ADJECTIVE - прилагательное
NOUN - существительное
VERB - глагол
NOUN - наречие
*/

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
		madLib <адрес файла.txt>`)
	} else if len(os.Args) == 2 {
		// проверка корректности имени файла и его открытие
		file, err := ioutil.ReadFile(os.Args[1])
		check(err)
		strFile := string(file)
		wordsArr := strings.Split(strFile, " ")
		for i := 0; i < len(wordsArr); i++ {
			fmt.Println(wordsArr[i])
		}
	}
}

// проверка на ошибки
func check(e error) {
	if e != nil {
		panic(e)
	}
}
