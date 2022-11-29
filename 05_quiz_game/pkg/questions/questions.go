package questions

import "fmt"

func AskQuestions(questions [][]string, score *int) {
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
