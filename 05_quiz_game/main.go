package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// the user can specify the csv file, other ./problems.csv will be the default
	// to add a csv use the --file flag when running the program
	var fileFlag string
	flag.StringVar(&fileFlag, "file", "", "to add a custom csv file")
	flag.Parse()

	if fileFlag == "" {
		fileFlag = "./problems.csv"
	}

	timer := time.NewTimer(5 * time.Second)
	done := make(chan bool)
	go func() {
		<-timer.C
		done <- true
	}()

	var score int

	go func() {
		//	records score

		questions, err := readCsvFile(fileFlag)
		if err != nil {
			return
		}

		askQuestions(questions, &score)
		fmt.Printf("You scored %d\n", score)
	}()

	if <-done {
		fmt.Println()
		fmt.Printf("time's up. Score: %d\n", score)
	}
}

func readCsvFile(file string) ([][]string, error) {
	var records [][]string

	// open csv file
	f, err := os.Open(file)

	if err != nil {
		return records, err
	}
	defer f.Close()

	//initialize new csv reader
	csvReader := csv.NewReader(f)

	records, err = csvReader.ReadAll()
	if err != nil {
		return records, err
	}
	return records, nil
}

func askQuestions(questions [][]string, score *int) {
	for index, value := range questions {
		var answer string
		fmt.Printf("Problem #%d: %s = ", index+1, value[0])
		fmt.Scan(&answer)
		fmt.Println()
		if answer == value[1] {
			*score += 1
		}
	}
}
