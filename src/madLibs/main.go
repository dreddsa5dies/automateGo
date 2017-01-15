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
		ADVERB - наречие
		Использование:
		madLib <адрес файла>`)
	} else if len(os.Args) == 2 {
		// проверка корректности имени файла и его открытие
		file, err := ioutil.ReadFile(os.Args[1])
		check(err)
		// формируем строку из []byte
		strFile := string(file)
		// разбиваем строку на массив слов
		// однако . относятся к слову (((
		wordsArr := strings.Split(strFile, " ")
		for i := 0; i < len(wordsArr); i++ {
			// кейсы по замене
			switch {
			case wordsArr[i] == "ADJECTIVE":
				fmt.Println("прилагательное")
			case wordsArr[i] == "NOUN":
				fmt.Println("существительное")
			case wordsArr[i] == "VERB":
				fmt.Println("глагол")
			case wordsArr[i] == "ADVERB":
				fmt.Println("наречие")
			}
		}
	}
}

// проверка на ошибки
func check(e error) {
	if e != nil {
		panic(e)
	}
}
