package view

import (
	"bufio"
	"fmt"
	"os"
)

func PromptUser() {
	fmt.Printf("Type in a city for it's weather > ")
}

//
//func PromptUserAgain() {
//	fmt.Printf("Type in a city for it's weather or press enter to exit > ")
//	reader := bufio.NewReader(os.Stdin)
//	response, _ := reader.ReadString('\n')
//	if len(response) == 1 {
//		os.Exit(2)
//	}
//
//}

func GetCity() string {
	reader := bufio.NewReader(os.Stdin)
	city, _ := reader.ReadString('\n')
	city = city[:len(city)-1]
	return city
}
