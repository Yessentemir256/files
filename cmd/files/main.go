package main

import (
	"io"
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

	content := make([]byte, 0)
	buf := make([]byte, 4)
	for {
		read, err := file.Read(buf)
		if err == io.EOF { // файл закончился
			content = append(content, buf[:read]...)
			break
		}

		if err != nil {
			log.Print(err)
			return
		}

		content = append(content, buf[:read]...)
	}
	data := string(content)
	log.Print(data)
}
