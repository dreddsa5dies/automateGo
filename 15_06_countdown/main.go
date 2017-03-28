// отсчет времени и exec
package main

import (
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	timeLeft := 2

	for timeLeft > 0 {
		println(timeLeft)
		time.Sleep(time.Second * 1)
		timeLeft--
	}

	// поиск исполняемого файла
	binary, lookErr := exec.LookPath("mpg123")
	if lookErr != nil {
		panic(lookErr)
	}

	// аргументы, обязательно вызываемый бинарник указывать
	args := []string{"mpg123", "alarm.wav"}

	env := os.Environ()

	// вызов бинарника
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
