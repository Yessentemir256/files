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
	defer func() {
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	log.Printf("%#v", file)
}
