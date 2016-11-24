package main

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// данные
	capitals := map[string]string{"Alabama": "Montgomery", "Alaska": "Juneau", "Arizona": "Phoenix", "Arkansas": "Littl Rock", "California": "Sacramento", "Colorado": "Denver", "Connecticut": "Hartford", "Delaware": "Dover", "Florida": "Tallahassee", "Georgia": "Atlanta", "Hawaii": "Honolulu", "Idaho": "Boise", "Illinois": "Springfield", "Indiana": "Indianapolis", "Iowa": "Des Moines", "Kansas": "Topeka", "Kentucky": "Frankfort", "Louisiana": "Baton Rouge", "Maine": "Augusta", "Maryland": "Annapolis", "Massachusetts": "Boston", "Michigan": "Lansing", "Minnesota": "Saint Paul", "Mississippi": "Jackson", "Missouri": "Jefferson City", "Montana": "Helena", "Nebraska": "Lincoln", "Nevada": "Carson City", "New Hampshire": "Concord", "New Jersey": "Trenton", "New Mexico": "Santa Fe", "New York": "Albany", "North Carolina": "Raleigh", "North Dakota": "Bismarck", "Ohio": "Columbus", "Oklahoma": "Oklahoma City", "Oregon": "Salem", "Pennsylvania": "Harrisburg", "Rhode Island": "Providence", "South Carolina": "Columbia", "South Dakota": "Pierre", "Tennessee": "Nashville", "Texas": "Austin", "Utah": "Salt Lake City", "Vermont": "Montpelier", "Virginia": "Richmond", "Washington": "Olympia", "West Virginia": "Charleston", "Wisconsin": "Madison", "Wyoming": "Cheyenne"}

	var capitalsItems, states, wrongAnswers, wrongAnswers2 []string

	for _, x := range capitals {
		capitalsItems = append(capitalsItems, x)
	}

	// генерация файлов
	for i := 1; i <= 35; i++ {
		// создание билетов и ключей
		str1 := "capitalsquiz" + string(i) + ".txt"
		quizFile, err := os.Create(str1)
		check(err)
		defer quizFile.Close()

		str2 := "capitalsquiz_answers" + string(i) + ".txt"
		answerKeyFile, err := os.Create(str2)
		check(err)
		defer answerKeyFile.Close()

		// запись заголовков
		quizFile.WriteString("Name:\n\nDate:\n\nPeriod:\n\n")
		str3 := "State Capitals Quiz (Form" + string(i) + ")"
		quizFile.WriteString(strings.Repeat(" ", 20) + str3)
		quizFile.WriteString("\n\n")

		// добавление столиц
		for x, _ := range capitals {
			states = append(states, x)
		}

		rand.Seed(time.Now().UnixNano())
		// перемешивание порядка следования столиц
		shuffle(states)

		// организация цикла по 50 штатам
		for i := 0; i < 50; i++ {
			// получение правильных и неправильных ответов
			correctAnswer := capitals[states[i]]
			for _, x := range capitals {
				wrongAnswers = append(wrongAnswers, x)
			}
			delete(wrongAnswers, correctAnswer)
			for j := 0; j < 3; j++ {
				wrongAnswers2 = append(wrongAnswers2, wrongAnswers[rand.Intn(len(wrongAnswers))])
			}
			answerOptions := make([]string, 4)
			copy(answerOptions, wrongAnswers2)
			answerOptions = append(answerOptions, correctAnswer)
			shuffle(answerOptions)

			// запись в файлы
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
