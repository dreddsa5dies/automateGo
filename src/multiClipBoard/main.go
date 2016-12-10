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
		mcb list			просмотр всех ключей
		mcb del				удаляет весь список и все данные`)
	} else {
		// проверка наличия файла
		// открываем с опциями добавления, записи и
		// создания если не существует
		mcb, _ := os.OpenFile("/tmp/mcb", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		defer mcb.Close()

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
			defer file.Close()
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
				// ветка удаления
			} else if os.Args[1] == "del" {
				// удалить все
				os.Remove("/tmp/mcb")
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
