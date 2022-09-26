package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type problem struct {
	Question string
	Answer   string
}

func main() {

	readCSV()

}

func readCSV() {
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		prob := problem{
			Question: line[0],
			Answer:   line[1],
		}
		fmt.Println(prob.Question + " " + prob.Answer)
	}
}
