package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	readCSV()
	//TODO call askValidate() from here instead

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
