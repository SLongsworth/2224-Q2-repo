package main

import (
	"fmt"
	"time"
)

func displayAnimals(name string) {

	fmt.Println(name)

}

func main() {
	//call goroutine
	go displayAnimals("Cow")
	displayAnimals("cat")
	go displayAnimals("Dog")
	go displayAnimals("Snake")
	go displayAnimals("sheep")
	displayAnimals("chimp")

	time.Sleep(2 * time.Second)

}
