package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {
	pwdDir, _ := os.Getwd()
	os.Mkdir(pwdDir+"/log", 0740)
	fLog, err := os.OpenFile(pwdDir+"/log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	check(err, fLog)

	excelFileName := pwdDir + "/example.xlsx"

	/*
		xlsx, err := excelize.OpenFile(excelFileName)
		check(err, fLog)

		// получение столбцов
		// не йас однако - 1 столбец неверно считывает
		x := xlsx.GetRows("Sheet1")
		fmt.Println(x)

		// имя страницы
		sheet1 := xlsx.GetSheetName(1)
		fmt.Println(sheet1)

		// получение ячейки
		cell := xlsx.GetCellValue("Sheet1", "A2")
		fmt.Println(cell)

		// только 1й столбец
		sheet := xlsx.GetRows(xlsx.GetSheetName(1))
		for _, row := range sheet {
			fmt.Println(row[0])
		}
	*/

	xlFile, err := xlsx.OpenFile(excelFileName)
	check(err, fLog)

	// Имя страницы
	i := xlFile.Sheets[0]
	fmt.Println(i.Name)

	// получение ячейки
	y := i.Cell(0, 0)
	a, _ := y.String()
	fmt.Println(a)

	// вывод 2 столбца
	// сколько всего строк
	k := i.MaxRow
	// до максимума строк пройти
	for u := 0; u < k; u++ {
		// по 2му столбцу
		sd := i.Cell(u, 1)
		// и вывести как строку
		ss, _ := sd.String()
		fmt.Println(ss)
	}
}

// err check to log
func check(err error, fLog *os.File) {
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))
}
