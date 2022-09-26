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
	csvFilename := flag.String("csv", "problems.csv", "gimme a CSV file but the format of 'question, answer'")
	flag.Parse()
	readCSV(*csvFilename)
	//TODO call askValidate() from here instead

}

func readCSV(csvFilename string) {
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
	for _, line := range csvLines {
		prob := Problem{
			Question: line[0],
			Answer:   line[1],
		}
		askValidate(prob)

	}
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
