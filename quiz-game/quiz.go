package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	csvName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer' (default 'problems.csv')")
	flag.Parse()
	csvFile, err := os.Open(*csvName)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvFile)
	problemCount, correctCount, userInput := 0, 0, ""
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		problem, solution := record[0], record[1]
		problemCount++
		fmt.Printf("Problem #%d: %s = ", problemCount, problem)
		_, err = fmt.Scan(&userInput)
		if err != nil {
			log.Fatal(err)
		}
		if solution == strings.Trim(userInput, " ") {
			correctCount++
		}
	}
	fmt.Printf("Total problems asked : %d \n", problemCount)
	fmt.Printf("Total problems solved: %d \n", correctCount)
}
