package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "gimme a CSV file with the format of 'question, answer'")
	flag.Parse()

	csvLines := readCSV(*csvFilename)

	for _, line := range csvLines {
		prob := Problem{
			Question: line[0],
			Answer:   line[1],
		}
		askValidate(prob)

	}

}

func readCSV(csvFilename string) [][]string {
	csvFile, err := os.Open(csvFilename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	return csvLines

}

func askValidate(class Problem) {
	var userAns string

	fmt.Println(class.Question, "Enter your ans: ")
	fmt.Scanln(&userAns)

	if userAns == class.Answer {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect")
	}

}
