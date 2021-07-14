package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type question struct {
	questionContent string
	answer          string
}

func main() {
	limit := flag.Int("limit", 30, "Time limit (seconds)")
	csv := flag.String("csv", "questions/questions.csv", "Questionnaire file path")

	flag.Parse()

	if *limit < 30 {
		fmt.Println("Limit cannot be less than 30 seconds!")
		os.Exit(1)
	}

	questionRecords := readQuestionsCSV(*csv)
	reader := bufio.NewReader(os.Stdin)
	rightAnswer := 0

	fmt.Println("--------QUIZ--------")
	fmt.Printf("You have %v seconds.\n", *limit)
	go timer(*limit)

	//fmt.Println(<-timeoutChan)

	for _, row := range questionRecords {
		q := new(question)
		q.questionContent = row[0]
		q.answer = row[1]

		fmt.Printf("Question: What is %v. Answer: ", q.questionContent)

		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimRight(userAnswer, "\r\n")
		if strings.Compare(q.answer, userAnswer) == 0 {
			rightAnswer = rightAnswer + 1
		}
	}
	fmt.Println("--------DONE--------")
	questionsLength := len(questionRecords)
	fmt.Printf("Score: %v/%v", rightAnswer, questionsLength)
}

func readQuestionsCSV(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}

func timer(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
	fmt.Printf("\nSorry, you ran out of time!")
	os.Exit(1)
}
