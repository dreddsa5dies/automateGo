// параллелизма тест
package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	i := randomInt()

	messages := make(chan int)

	var wg sync.WaitGroup

	// обязательно ожидание всех горутин
	wg.Add(len(i))

	// тут в одно время рандомно обработка
	for _, arr := range i {
		go func(arr int) {
			defer wg.Done()
			time.Sleep(time.Second * 2)
			messages <- arr
		}(arr)
		// ввод и вывод значения
	}

	// тут вывод по порядку т.к. задержка равна значению
	// for _, arr := range i {
	// 	go func(arr int) {
	// 		defer wg.Done()
	// 		time.Sleep(time.Second * time.Duration(arr))
	// 		messages <- arr
	// 	}(arr)
	// }

	go func() {
		for j := range messages {
			println(j)
		}
	}()

	wg.Wait()
}

func randomInt() []int {
	list := rand.Perm(20)
	for i := 0; i < len(list); i++ {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
	return list
}
