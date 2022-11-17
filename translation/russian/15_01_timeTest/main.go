// задержка расчет времени округление и большие числа, дата их сравнение и форматирование
package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	// время начала
	startTime := time.Now()
	prod := calcTime()

	// время конца
	endTime := time.Now()

	ti := round(endTime.Sub(startTime).Seconds(), 3)

	//приведение большого Int к строке и вычисление ее длины
	fmt.Printf("Длина результата %v\n", len(prod.String()))

	// endTime.Sub(startTime).Seconds() - из конца вычитаем начало в формате секунд
	fmt.Printf("Расчет по времени %1.3f сек.\n", ti)

	// test задержки
	for i := 0; i < 5; i++ {
		fmt.Println("Тик")
		time.Sleep(1 * time.Second)
		fmt.Println("Так")
		time.Sleep(1 * time.Second)
	}

	// получение даты
	year, mounth, day := time.Now().Date()
	fmt.Println(year, mounth.String(), day)

	// сравнение указанной даты и текущей
	halloween2017 := time.Date(2017, 10, 31, 0, 0, 0, 0, time.Local)
	// до тех пор пока Хэллоуин не наступит
	for halloween2017.Before(time.Now()) {
		fmt.Println("ok")
		time.Sleep(time.Second * 1)
	}

	// преобразование даты в строки
	println(halloween2017.Format("2006-01-02"))
	// 01 - месяц, 02 - день, 2006 - год
	println(halloween2017.Format("01.02.2006"))

	// преобразование строки в дату
	dataTime := "2015.03.11"
	dateParse, _ := time.Parse("2006.01.02", dataTime)
	fmt.Println(dateParse)
	fmt.Println(dateParse.Month().String())
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

// функция округления
func round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
