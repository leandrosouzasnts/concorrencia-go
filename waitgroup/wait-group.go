package main

import (
	"fmt"
	"sync"
	"time"
)

//Sincronismo de goroutines (não é tão utilizado)
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

	waitGroup.Wait() //0

}

func escrever(text string) {

	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
