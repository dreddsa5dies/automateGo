// сохраняет и загружает фрагменты текста
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	if len(os.Args) == 1 {
		// вывод справки
		fmt.Println(`Ипользование:
		mcb save <ключевое слово>	сохраняет буфер по ключевому слову
		mcb <ключевое слово>		загружает данные по ключевому слову в буфер
		mcb list			view all key-words
		mcb del <ключевое слово>	удаляет данные из списка
		mcb del				удаляет весь список и все данные`)
	} else {
		// проверка наличия файла
		// открываем с опциями добавления, записи и
		// создания если не существует
		mcb, _ := os.OpenFile("/tmp/mcb", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

		mapBase := make(map[string]string)
		// парсинг флагов
		if len(os.Args) == 3 && os.Args[1] == "save" {
			// запись в переменную из буфера
			text, _ := clipboard.ReadAll()
			mapBase[os.Args[2]] = text
			mapB, _ := json.Marshal(mapBase)
			// добавляю символ новой строки
			newStr := "\n"
			mapB = append(mapB, newStr...)
			// запись в файл хранения
			if _, err := mcb.Write(mapB); err != nil {
				panic(err)
			}
		} else if len(os.Args) == 2 {
			// считываем файл по строкам
			// записываем в мапу
			file, _ := os.Open("/tmp/mcb")
			scanner := bufio.NewScanner(file)
			mapStr := make(map[string]string)
			for scanner.Scan() {
				j := scanner.Text()
				err := json.Unmarshal([]byte(j), &mapStr)
				check(err)
			}
			// вывод списка ключей
			if os.Args[1] == "list" {
				fmt.Println("Ключевые слова:")
				for x, _ := range mapStr {
					fmt.Println("\t" + x)
				}
			} else if os.Args[1] == "del" {
				var strDel string
				fmt.Println("Удалить все?\tYes/No")
				fmt.Scan(&strDel)
				if strDel == "Yes" {
					os.Remove("/tmp/mcb")
				} else if strDel == "No" {
					fmt.Println("How key string?")
				} else {
					fmt.Println("No key")
				}
			} else {
				// вставка в буфер по ключу
				for x, y := range mapStr {
					if os.Args[1] == x {
						clipboard.WriteAll(y)
					}
				}
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
