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
		madLib <адрес файла.txt>`)
	} else if len(os.Args) == 2 {
		// проверка корректности имени файла и его открытие
		file, err := os.Open(os.Args[1])
		check(err)
		defer file.Close()
		fmt.Println(file)
	}
}

// проверка на ошибки
func check(e error) {
	if e != nil {
		panic(e)
	}
}
