// отсчет времени и exec
package main

import (
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

func main() {
	timeLeft := 2

	for timeLeft > 0 {
		print(strconv.Itoa(timeLeft) + " ")
		time.Sleep(time.Second * 1)
		timeLeft--
	}
	print("\n")

	// поиск исполняемого файла
	binary, lookErr := exec.LookPath("cvlc")
	if lookErr != nil {
		panic(lookErr)
	}

	// аргументы, обязательно вызываемый бинарник указывать
	args := []string{"cvlc", "alarm.wav"}

	env := os.Environ()

	// вызов бинарника
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
