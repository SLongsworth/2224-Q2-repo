//Creating channels

package main

import (
	"fmt"
)

func main() {
	//Create a channel in the main function
	riddleQuestion := make(chan string)
	riddleAnswer := make(chan string)

	go riddleData(riddleQuestion, riddleAnswer)

	fmt.Println(<-riddleQuestion)
	fmt.Println(<-riddleAnswer)

}

func riddleData(riddleQuestion chan string, riddleAnswer chan string) {
	riddleQuestion <- "What becomes smaller when you turn it upside down?"
	riddleAnswer <- "The number 9!"
}
