package main

import (
	"fmt"
	"sync"
)

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	var waitGroup sync.WaitGroup

	numberWorks := 6
	for i := 0; i < numberWorks; i++ {
		waitGroup.Add(1)
		go worker(&waitGroup, tarefas, resultados)
	}

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	go func() {
		waitGroup.Wait()
		close(resultados)
	}()

	for resultado := range resultados {
		fmt.Println(resultado)
	}
}

func worker(waitGroup *sync.WaitGroup, tarefas <-chan int, resultados chan<- int) {
	defer waitGroup.Done()
	for numero := range tarefas {
		resultados <- int(fibonaci(numero))
	}
}

func fibonaci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonaci(posicao-2) + fibonaci(posicao-1)
}
