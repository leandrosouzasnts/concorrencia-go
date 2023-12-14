package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexador(escrever("Olá Leandro"), escrever("Olá Divani"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexador(canal1, canal2 <-chan string) <-chan string {
	output := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				output <- mensagem
			case mensagem := <-canal2:
				output <- mensagem
			}
		}
	}()
	return output
}

func escrever(texto string) <-chan string { //Generator é usado para abstrair uma goroutine e retornamos um canal
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Int63n(2000)))
		}
	}()

	return canal
}
