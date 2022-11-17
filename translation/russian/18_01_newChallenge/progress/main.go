package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/schollz/progressbar"
)

func main() {
	// cli
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %v IntDigital\n", os.Args[0])
		os.Exit(1)
	} else {
		// n - digit for progress bar
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			// error for stupid users
			log.Fatalln("Not Int Digital", err)
		}

		// default example
		bar := progressbar.New(n)
		for i := 0; i < n; i++ {
			bar.Add(1)
			time.Sleep(10 * time.Millisecond)
		}
		os.Exit(0)
	}
}
