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
			// проверка на вхождение в слово точки
			if strings.Contains(wordsArr[i], ".") {
				wordsArr[i] = wordsArr[i][:len(wordsArr[i])-1]
			}
			// кейсы по замене
			switch {
			case wordsArr[i] == "ADJECTIVE":
				fmt.Println("Введите имя прилагательное:")
				s := scan()
				wordsArr[i] = s
			case wordsArr[i] == "NOUN":
				fmt.Println("Введите имя существительное:")
				s := scan()
				wordsArr[i] = s
			case wordsArr[i] == "VERB":
				fmt.Println("Введите глагол:")
				s := scan()
				wordsArr[i] = s
			case wordsArr[i] == "ADVERB":
				fmt.Println("Введите наречие:")
				s := scan()
				wordsArr[i] = s
			}
		}
		fmt.Println(wordsArr)
	}
}

// проверка на ошибки
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// функция ввода данных
func scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
