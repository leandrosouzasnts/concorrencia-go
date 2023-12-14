package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("1")
	canal2 := escrever("2")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
		fmt.Println(<-canal2)
	}
}

func escrever(texto string) <-chan string { //Generator é usado para abstrair uma goroutine e retornamos um canal
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 5000)
		}
	}()

	return canal
}
