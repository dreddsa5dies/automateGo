// поиск и замена слов в txt
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
		reference := `
ИМЯ

madLibs -    Замена ключевых слов в файле
			ADJECTIVE - прилагательное
			NOUN - существительное
			VERB - глагол
			ADVERB - наречие
			Использование:
			madLib <адрес файла> 

СИНТАКСИС

madLibs ФАЙЛ  

АВТОР

Виктор Соловьев 

СООБЩЕНИЕ ОБ ОШИБКАХ

Об ошибках сообщайте по адресу <viktor.vladimirovich.solovev@gmail.com>.  

АВТОРСКИЕ ПРАВА

Copyright 2017 Viktor Solovev

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.`
		fmt.Println(reference)
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
