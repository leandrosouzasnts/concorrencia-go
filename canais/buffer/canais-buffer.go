package main

import "fmt"

func main() {

	canal := make(chan string, 2)

	canal <- "Olá Mundo"
	canal <- "Olá Universo"

	mensagem := <-canal
	fmt.Println(mensagem)

}
