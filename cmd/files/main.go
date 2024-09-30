package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("data/readme.txt")
	if err != nil {
		log.Print(err)
		return
	}
	defer closeFile(file) // будет выполнено когда выйдем из main
	log.Printf("%#v", file)
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Print(err)
	}
}
