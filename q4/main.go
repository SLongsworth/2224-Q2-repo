package main

import (
	"fmt"
	"sync"
)

func sunflower(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("sunflowers")
}

func rose(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("roses")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go sunflower(&wg)
	go rose(&wg)

	//delay
	wg.Wait()
	fmt.Println("Those are my favorite flowers")
}
