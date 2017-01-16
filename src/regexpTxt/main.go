package main

/*
поиск в *.txt по регулярному выражению
*/

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		// вывод справки
		fmt.Println(`поиск во всех *.txt в папке по регулярному выражению
		Использование:
		regexpTxt <адрес папки>`)
	} else if len(os.Args) == 2 {
		//открытие папки
		dir, err := os.Open(os.Args[1])
		if err != nil {
			return
		}
		defer dir.Close()

		// список файлов
		fileInfos, err := dir.Readdir(-1)
		if err != nil {
			return
		}
		// создание шаблона регулярного выражения
		fmt.Print("Регулярное выражение> ")
		regexTxt, _ := regexp.Compile(scan())

		for _, fi := range fileInfos {
			// проверка на .txt
			if strings.HasSuffix(fi.Name(), ".txt") {
				// проверка корректности имени файла и его открытие
				file, err := ioutil.ReadFile(fi.Name())
				if err != nil {
					fmt.Println(err)
				}
				// формируем строку из []byte
				strFile := string(file)
				// поиск по регулярному выражению
				var founds []string
				if regexTxt.MatchString(strFile) {
					founds = regexTxt.FindAllString(strFile, -1)
				}

				// вывод данных
				if len(founds) > 0 {
					fmt.Printf("Founds in %v:\n", fi.Name())
					fmt.Println(strings.Join(founds, "\n"))
				} else {
					fmt.Println("Not found.")
				}
			}
		}
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
