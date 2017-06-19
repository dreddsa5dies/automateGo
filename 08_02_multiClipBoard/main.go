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
		reference := `
ИМЯ

mcb - многостраничный буфер обмена (для работы с буфером требуется xclip)

СИНТАКСИС

mcb save <ключевое слово>	сохраняет буфер по ключевому слову
mcb <ключевое слово>		загружает данные по ключевому слову в буфер
mcb list			просмотр всех ключей
mcb del <ключевое слово>	удаляет буфер по ключевому слову
mcb del				удаляет весь список и все данные
`
		fmt.Println(reference)
	} else {
		// проверка наличия файла
		// открываем с опциями добавления, записи и
		// создания если не существует
		mcb, _ := os.OpenFile("/tmp/mcb", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		defer mcb.Close()

		newStr := "\n"
		mapBase := make(map[string]string)
		// парсинг флагов
		if len(os.Args) == 3 && os.Args[1] == "save" {
			// запись в переменную из буфера
			text, _ := clipboard.ReadAll()
			mapBase[os.Args[2]] = text
			mapB, _ := json.Marshal(mapBase)
			// добавляю символ новой строки
			mapB = append(mapB, newStr...)
			// запись в файл хранения
			if _, err := mcb.Write(mapB); err != nil {
				panic(err)
			}
		} else if len(os.Args) >= 2 {
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
			// удаление по ключу
			if len(os.Args) == 3 && os.Args[1] == "del" {
				// создание временного файла
				mcbTemp, _ := os.OpenFile("/tmp/.mcbTemp", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				defer mcbTemp.Close()
				// поиск и удаление по ключу
				for x, _ := range mapStr {
					if x == os.Args[2] {
						delete(mapStr, x)
					}
				}
				// создание и запись новой структуры
				// кривовато причем
				mapStr, _ := json.Marshal(mapStr)
				mapStr = append(mapStr, newStr...)
				// запись в файл хранения
				if _, err := mcbTemp.Write(mapStr); err != nil {
					panic(err)
				}
				os.Rename("/tmp/.mcbTemp", "/tmp/mcb")
				os.Remove("/tmp/.mcbTemp")
			}
			// вывод списка ключей
			if os.Args[1] == "list" {
				fmt.Println("Ключевые слова:")
				for x, _ := range mapStr {
					fmt.Println("\t" + x)
				}
				// ветка удаления
			} else if len(os.Args) == 2 && os.Args[1] == "del" {
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
