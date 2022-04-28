package view

import (
	"bufio"
	"fmt"
	"os"
)

func PromptUser() {
	fmt.Printf("Type in a city for it's weather > ")
}

func GetCity() string {
	reader := bufio.NewReader(os.Stdin)
	city, _ := reader.ReadString('\n')
	city = city[:len(city)-1]
	return city
}
