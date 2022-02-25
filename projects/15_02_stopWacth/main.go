// секундометр
package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup() {
	fmt.Println("\nГотово")
}

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// перехват CTRL+C
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()

	secundometr()
}

// функция секундометр
func secundometr() {
	println("Чтобы начать отсчет, нажмите ENTER.\nВпоследствии для имитации щелчков кнопки секундомера нажимайте клавишу ENTER.\nДля выхода из програмы нажмите клавиши CTRL+C")

	fmt.Scanln()
	println("Отсчет начат")

	startTime := time.Now()
	lastTime := startTime
	lapNum := 1

	for {
		fmt.Scanln()
		lapTime := round(time.Now().Sub(lastTime).Seconds(), 2)
		totalTime := round(time.Now().Sub(startTime).Seconds(), 2)
		fmt.Printf("Замер №%v:\tвсего прошло %v\tс последнего замера %v\n", lapNum, totalTime, lapTime)
		lapNum++
		lastTime = time.Now()
	}
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
