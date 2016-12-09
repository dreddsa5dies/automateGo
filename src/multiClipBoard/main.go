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
		mcb del				удаляет весь список и все данные
		mcb del и далее <ключевое слово>	удаляет по ключевому слову`)
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
				// ветка удаления
			} else if os.Args[1] == "del" {
				var strDel string
				fmt.Println("Удалить все?\tYes/No")
				fmt.Scan(&strDel)
				// если удалить все
				if strDel == "Yes" {
					os.Remove("/tmp/mcb")
				} else if strDel == "No" {
					// TODO добавить изменение файла
					var keyDel string
					fmt.Print("Удаление ключевого слова ")
					fmt.Scan(&keyDel)
					// открытие файла для изменения
					mcbFile, _ := os.OpenFile("/tmp/mcb", os.O_WRONLY, 0600)
					defer mcbFile.Close()
					// считывание всех данных
					neWscanner := bufio.NewScanner(mcbFile)
					mapStrFile := make(map[string]string)
					for neWscanner.Scan() {
						p := scanner.Text()
						err := json.Unmarshal([]byte(p), &mapStrFile)
						check(err)
					}
					for x, _ := range mapStrFile {
						if x == keyDel {
							delete(mapStrFile, x)
						}
					}
					mapBas, _ := json.Marshal(mapStrFile)
					// добавляю символ новой строки
					newStr1 := "\n"
					mapBas = append(mapBas, newStr1...)
					// запись в файл хранения
					if _, err := mcbFile.Write(mapBas); err != nil {
						panic(err)
					}
				} else {
					fmt.Println("Нечего удалять")
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
