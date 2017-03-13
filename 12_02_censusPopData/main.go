package main

import (
	"log"
	"os"

	"fmt"

	flags "github.com/jessevdk/go-flags"
	"github.com/tealeg/xlsx"
)

var opts struct {
	FileExel string `short:"o" long:"open" default:"censuspopdata.xlsx" description:"Региональная численность населения"`
}

func main() {
	// разбор флагов
	flags.Parse(&opts)

	// в какой папке исполняемы файл
	pwdDir, _ := os.Getwd()

	// создание файла log для записи ошибок
	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	// запись ошибок и инфы в файл
	log.SetOutput(fLog)

	excelFileName := pwdDir + "/" + opts.FileExel

	log.Printf("Открытие рабочей книги: %v", excelFileName)

	wb, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("Ошибка открытия книги %v", err)
	}

	sheet := wb.Sheet["Population by Census Tract"]
	log.Printf("Открытие страницы: %v", sheet.Name)

	// сколько всего строк
	k := sheet.MaxRow

	// словарь для записи
	type Key struct {
		State, Country string
	}
	popdata := make(map[Key]float64)

	// до максимума строк пройти
	for u := 1; u < k; u++ {
		// получение данных ячеек формата *Cell
		state := sheet.Cell(u, 1)
		country := sheet.Cell(u, 2)
		pop := sheet.Cell(u, 3)

		// перевод *Cell  в нормальные значения
		nameStateStr, err := state.String()
		if err != nil {
			log.Fatalf("Ошибка nameStateStr %v", err)
		}
		countryStr, err := country.String()
		if err != nil {
			log.Fatalf("Ошибка countryStr %v", err)
		}
		popFl, err := pop.Float()
		if err != nil {
			log.Fatalf("Ошибка popFl %v", err)
		}

		// количество населения по районам
		popdata[Key{nameStateStr, countryStr}] += popFl
	}

	fmt.Println(popdata[Key{"WY", "Sheridan"}])
}
