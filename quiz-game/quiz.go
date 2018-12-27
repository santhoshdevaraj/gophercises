package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var questionNumber int

// createMapFromFile converts the CSV file into a map of questions and solution
func createMapFromFile(csvName *string) map[string]string {
	csvFile, err := os.Open(*csvName)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvFile)
	questions := map[string]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		questions[record[0]] = record[1]
	}
	return questions
}

func askQuestion(q string, channel chan string) {
	var userInput string
	questionNumber++
	fmt.Printf("Problem #%d: %s = ", questionNumber, q)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		log.Fatal(err)
	}
	channel <- userInput
}

func main() {
	csvName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	limit := flag.Float64("limit", 30, "the time limit for the quiz in seconds (default 30)")
	flag.Parse()
	questions := createMapFromFile(csvName)
	startTime := time.Now()
	channel := make(chan string)
	answerCount := 0

	for question := range questions {
		remTime := int((*limit)*60 - time.Since(startTime).Seconds())
		go askQuestion(question, channel)
		select {
		case userInput := <-channel:
			if strings.Trim(userInput, " ") == questions[question] {
				answerCount++
			}
		case <-time.After(time.Duration(remTime) * time.Second):
			fmt.Println("\nAllotted time has completed. Quiz will end now")
			break
		}
	}
	fmt.Printf("\nquestions asked : %d", len(questions))
	fmt.Printf("\ncorrect answers : %d", answerCount)
}
