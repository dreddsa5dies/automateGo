package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strconv"
)

type Task struct {
	Content  string
	Complete bool
}

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Usage = "add, list, and complete tasks"
	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "add a task",
			Action: func(c *cli.Context) {
				task := Task{Content: c.Args().First(), Complete: false}
				AddTask(task, "tasks.json")
				fmt.Println("added task: ", c.Args().First())
			},
		},
		{
			Name:  "complete",
			Usage: "complete a task",
			Action: func(c *cli.Context) {
				idx, err := strconv.Atoi(c.Args().First())
				if err != nil {
					panic(err)
				}
				CompleteTask(idx)
				fmt.Println("completed task: ", c.Args().First())
			},
		},
		{
			Name:      "list",
			ShortName: "ls",
			Usage:     "print all uncompleted tasks in list",
			Action: func(c *cli.Context) {
				ListTasks()
			},
		},
	}
	app.Run(os.Args)
}

func OpenTaskFile() *os.File {
	// Проверяем существование файла
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		log.Fatal("tasks file does not exist")
		return nil
	}
	file, err := os.Open("tasks.json")
	if err != nil {
		panic(err)
	}
	return file
}

func AddTask(task Task, filename string) {
	j, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	// Добавляем перенос на новую строку
	// для лучшей читабельности
	j = append(j, "\n"...)
	// Open tasks.json in append-mode.
	f, _ := os.OpenFile(filename,
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	// Append our json to tasks.json
	if _, err = f.Write(j); err != nil {
		panic(err)
	}
}

func ListTasks() {
	file := OpenTaskFile()
	scanner := bufio.NewScanner(file)
	// Наш индекс, который мы будем использовать как номер задачи
	i := 1
	// `scanner.Scan()` перемещает сканер к следующему разделителю
	// и возвращает true. По умолчанию разделитель это перенос на новую
	//строку. Когда сканер доходит до конца файла, то возвращает false.
	for scanner.Scan() {
		// `scanner.Text()` возвращает текущий токен как строку
		j := scanner.Text()
		t := Task{}
		// Мы передаем в Unmarshall json строку конвертированную в байт слайс
		// и указатель на переменную типа `Task`. Поля этой переменной будут
		// заполнены значениями из json.
		err := json.Unmarshal([]byte(j), &t)
		// По умолчанию мы будем показывать только
		// не выполненные задания
		if err != nil {
			panic(err)
		}
		if !t.Complete {
			fmt.Printf("[%d] %s\n", i, t.Content)
			i++
		}
	}
}

func CompleteTask(idx int) {
	file := OpenTaskFile()
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		j := scanner.Text()
		t := Task{}
		err := json.Unmarshal([]byte(j), &t)
		if err != nil {
			panic(err)
		}
		if !t.Complete {
			if idx == i {
				t.Complete = true
			}
			i++
		}
		// Добавляем текущую задачу к временному файлу.
		// Обратите внимание, когда мы вызываем эту функцию
		// первый раз, то создается файл и записывается задача.
		AddTask(t, ".tempfile")
	}
	// Когда мы записали все в .tempfile, заменяем им файл tasks.json
	os.Rename(".tempfile", "tasks.json")
	// Теперь можем удаляять .tempfile
	os.Remove(".tempfile")
}
