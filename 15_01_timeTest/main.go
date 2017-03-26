//считывание JSON
package main

import (
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	// в какой папке исполняемый файл
	pwdDir, _ := os.Getwd()

	// создание файла log для записи ошибок
	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	// запись ошибок и инфы в файл
	log.SetOutput(fLog)

	// время начала
	startTime := time.Now()

	// время начала
	prod := calcTime()

	// время конца
	endTime := time.Now()

	//приведение большого Int к строке и вычисление ее длины
	log.Printf("Длина результата %v", len(prod.String()))

	// endTime.Sub(startTime).Seconds() - из конца вычитаем начало в формате секунд
	log.Printf("Расчет по времени %1.3f сек.", endTime.Sub(startTime).Seconds())
}

func calcTime() *big.Int {

	verybig := big.NewInt(1)
	ten := big.NewInt(10)

	// много раз 1 умножить на 10
	for i := 0; i < 100000; i++ {
		temp := new(big.Int)
		temp.Mul(verybig, ten)
		verybig = temp
	}

	return verybig
}
