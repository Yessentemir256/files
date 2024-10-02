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

	// читаем 4096 байт
	buf := make([]byte, 4096)
	read, err := file.Read(buf)
	if err != nil {
		log.Print(err)
		return
	}

	// сохраняем ровно столько, сколько прочитали
	data := string(buf[:read])
	log.Print(data)
}
