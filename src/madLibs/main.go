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
		var strDot, dotEntry []int
		for i := 0; i < len(wordsArr); i++ {
			// проверка на вхождение в слово точки
			if strings.Contains(wordsArr[i], ".") {
				strDot = append(strDot, i)
				wordsArr[i] = wordsArr[i][:len(wordsArr[i])-1]
			}
			// замена и контроль точки
			x, y := reInter(wordsArr[i])
			if y {
				dotEntry = append(dotEntry, i)
			}
			wordsArr[i] = x
		}
		// добавление потерянной точки
		for i := 0; i < len(strDot); i++ {
			for j := 0; j < len(dotEntry); j++ {
				if strDot[i] == dotEntry[j] {
					wordsArr[strDot[i]] = wordsArr[strDot[i]] + "."
				}
			}
		}
		// TODO: запись в файл
		// только для моего примера
		// если файл будет большой, то с переносами строк будет все плохо
		tmp, _ := os.OpenFile("tmp_"+os.Args[1], os.O_WRONLY|os.O_CREATE, 0600)
		defer tmp.Close()
		// []string to string
		var strAdd string
		for i := 0; i < len(wordsArr); i++ {
			strAdd += wordsArr[i]
			if i != len(wordsArr)-1 {
				strAdd += " "
			}
			if i == len(wordsArr)-1 {
				strAdd += "\n"
			}
		}
		// запись в файл хранения
		if _, err := tmp.Write([]byte(strAdd)); err != nil {
			panic(err)
		}
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

// функция замены слова
func reInter(strInter string) (string, bool) {
	var dotBool bool
	dotBool = false
	// кейсы по замене
	switch {
	case strInter == "ADJECTIVE":
		fmt.Println("Введите имя прилагательное:")
		s := scan()
		strInter = s
		dotBool = true
	case strInter == "NOUN":
		fmt.Println("Введите имя существительное:")
		s := scan()
		strInter = s
		dotBool = true
	case strInter == "VERB":
		fmt.Println("Введите глагол:")
		s := scan()
		strInter = s
		dotBool = true
	case strInter == "ADVERB":
		fmt.Println("Введите наречие:")
		s := scan()
		strInter = s
		dotBool = true
	}
	return strInter, dotBool
}
