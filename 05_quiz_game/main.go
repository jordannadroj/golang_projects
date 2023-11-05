package main

import (
	"flag"
	"fmt"
	qsts "github.com/jordannadroj/52_in_52/05_quiz_game/pkg/questions"
	"github.com/jordannadroj/52_in_52/05_quiz_game/pkg/read-csv"
	"time"
)

func main() {
	//Todo: Write Tests.

	// the user can specify the csv file, other ./problems.csv will be the default
	// to add a csv use the --file flag when running the program
	var (
		fileFlag     string
		timerFlag    int
		score        int
		numQuestions int
	)

	flag.StringVar(&fileFlag, "file", "./problems.csv", "to add a custom csv file containing a list of problems. If no file is added the default file 'problems.csv' file will be used")
	flag.IntVar(&timerFlag, "time", 20, "add a custom timer for the quiz. If not timer is specified a timer of 20 seconds will be the default")
	flag.Parse()
	done := make(chan bool)

	//	Read from CSV file and retrieve questions
	questions, err := read_csv.ReadCsvFile(fileFlag, &numQuestions)
	if err != nil {
		return
	}

	var start time.Time

	go func() {
		start = time.Now()
		qsts.AskQuestions(questions, &score)
		done <- true
	}()

	select {
	case <-time.After(time.Duration(timerFlag) * time.Second):
		fmt.Println()
		fmt.Printf("time's up. Score: %d/%d\n", score, numQuestions)
	case <-done:
		fmt.Printf("You scored %d/%d in %.2f seconds \n", score, numQuestions, time.Now().Sub(start).Seconds())
	}
}
