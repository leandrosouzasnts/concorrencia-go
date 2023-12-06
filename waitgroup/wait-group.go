package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Hello Word")
		waitGroup.Done() //-1
	}()

	go func() {
		escrever("Programmer GO")
		waitGroup.Done() //-1
	}()

	waitGroup.Wait()

}

func escrever(text string) {

	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
