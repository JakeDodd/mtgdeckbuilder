package main

import (
	//"fmt"
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("all-cards-20241022215316.json")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	start := time.Now()
	for true {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if line == "" {

		}
	}
	elapsed := time.Since(start)
	log.Println(elapsed)
}
