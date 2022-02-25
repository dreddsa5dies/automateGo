// генерация файлов
package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// данные
	capitals := map[string]string{
		"Alabama":        "Montgomery",
		"Alaska":         "Juneau",
		"Arizona":        "Phoenix",
		"Arkansas":       "Littl Rock",
		"California":     "Sacramento",
		"Colorado":       "Denver",
		"Connecticut":    "Hartford",
		"Delaware":       "Dover",
		"Florida":        "Tallahassee",
		"Georgia":        "Atlanta",
		"Hawaii":         "Honolulu",
		"Idaho":          "Boise",
		"Illinois":       "Springfield",
		"Indiana":        "Indianapolis",
		"Iowa":           "Des Moines",
		"Kansas":         "Topeka",
		"Kentucky":       "Frankfort",
		"Louisiana":      "Baton Rouge",
		"Maine":          "Augusta",
		"Maryland":       "Annapolis",
		"Massachusetts":  "Boston",
		"Michigan":       "Lansing",
		"Minnesota":      "Saint Paul",
		"Mississippi":    "Jackson",
		"Missouri":       "Jefferson City",
		"Montana":        "Helena",
		"Nebraska":       "Lincoln",
		"Nevada":         "Carson City",
		"New Hampshire":  "Concord",
		"New Jersey":     "Trenton",
		"New Mexico":     "Santa Fe",
		"New York":       "Albany",
		"North Carolina": "Raleigh",
		"North Dakota":   "Bismarck",
		"Ohio":           "Columbus",
		"Oklahoma":       "Oklahoma City",
		"Oregon":         "Salem",
		"Pennsylvania":   "Harrisburg",
		"Rhode Island":   "Providence",
		"South Carolina": "Columbia",
		"South Dakota":   "Pierre",
		"Tennessee":      "Nashville",
		"Texas":          "Austin",
		"Utah":           "Salt Lake City",
		"Vermont":        "Montpelier",
		"Virginia":       "Richmond",
		"Washington":     "Olympia",
		"West Virginia":  "Charleston",
		"Wisconsin":      "Madison",
		"Wyoming":        "Cheyenne",
	}

	// тут все столицы и штаты
	var states, capitalsItems []string
	for y, x := range capitals {
		capitalsItems = append(capitalsItems, x)
		states = append(states, y)
	}

	// генерация файлов
	for i := 0; i < 35; i++ {
		// создание билетов и ключей
		str1 := "билет_" + strconv.Itoa(i+1) + ".txt"
		quizFile, err := os.Create(str1)
		check(err)
		defer quizFile.Close()

		str2 := "билетОтвет_" + strconv.Itoa(i+1) + ".txt"
		answerKeyFile, err := os.Create(str2)
		check(err)
		defer answerKeyFile.Close()

		// запись заголовков
		quizFile.WriteString("ФИО:\n\nГруппа:\n\nДата:\n\n")
		str3 := "Билет " + strconv.Itoa(i+1)
		quizFile.WriteString(strings.Repeat(" ", 20) + str3)
		quizFile.WriteString("\n\n")

		rand.Seed(time.Now().UnixNano())
		// перемешивание порядка следования столиц
		shuffle(states)

		// организация цикла по 50 штатам
		for j := 0; j < 50; j++ {
			// получение правильных и неправильных ответов
			correctAnswer := capitals[states[j]]
			wrongAnswers := make([]string, len(capitalsItems))
			copy(wrongAnswers, capitalsItems)

			// удаление правильного ответа из неправильных
			answNoCorrect := make([]string, len(wrongAnswers)-1)
			for l := 0; l < len(wrongAnswers); l++ {
				if wrongAnswers[l] == correctAnswer {
					copy(answNoCorrect, removeAtIndex(wrongAnswers, l))
				}
			}

			// создание списка 3 случайных неправильных ответов
			// итоговые вариаты ответов
			var answerOptions []string
			for l := 0; l < 3; l++ {
				answerOptions = append(answerOptions, answNoCorrect[l])
			}
			answerOptions = append(answerOptions, correctAnswer)
			shuffle(answerOptions)

			// запись в файл
			str3 := strconv.Itoa(j+1) + " Выберите столицу штата " + states[j] + "\n"
			quizFile.WriteString(str3)
			strAbcd := "ABCD"
			for l := 0; l < 4; l++ {
				strAnsw := string(strAbcd[l]) + ". " + answerOptions[l] + "\n"
				quizFile.WriteString(strAnsw)
			}
			quizFile.WriteString("\n")

			// ключ ответа
			strAnswerOk := ""
			for l := 0; l < len(answerOptions); l++ {
				if answerOptions[l] == correctAnswer {
					strAnswerOk += string(strAbcd[l])
				}
			}
			strCorAnsw := strconv.Itoa(j+1) + ". " + strAnswerOk + "\n"
			answerKeyFile.WriteString(strCorAnsw)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func shuffle(a []string) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func removeAtIndex(source []string, index int) []string {
	lastIndex := len(source) - 1
	//меняем последнее значение и значение, которое хотим удалить, местами
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
