package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "customer_segmentation_data.csv"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the first 10 lines of the CSV file and print them to the console
	for i := 0; i < 10; i++ {
		line, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(line)
	}
}
