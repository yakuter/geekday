package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("geekday.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	args := os.Args[1:]

	for i := range args {
		file.WriteString(args[i] + "\n")
	}

}
