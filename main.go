package main

import (
	//"fmt"
	"bufio"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/JakeDodd/mtgdataload/models"
)

// TODO
// Create struct for original json card table: FileCard
// For each created table we need a matching struct: Card, Printcard, Langcard, etc
// Make new Folder in dataload project called models - models.go (package models at top) - all structs go in here

func main() {
	var card models.Card
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
		//line starts with [ and ends with ], we dont want to unmarshal the first and last line
		if line != "[\n" && line != "]" {
			//removing new line character, remove the comma in all but the last line, theres no comma in last line
			if line[len(line)-2:] == ",\n" {
				line = line[:len(line)-2]
			} else {
				line = line[:len(line)-1]
			}

			err = json.Unmarshal([]byte(line), &card)
			if err != nil {
				log.Println(line)
				log.Fatal(err)
			}
			log.Println(card.Name)
		}

	}
	elapsed := time.Since(start)
	log.Println(elapsed)
}
