package main

import (
	"fmt"
	"time"
)

func main() {

	chanel := make(chan string)

	go escrever("Olá mundo", chanel)

	fmt.Println("Iniciando programa!")
	for message := range chanel { //Fica verificando se o canal está aberto, quando fechado o mesmo vai finalizar o processo do for
		fmt.Println(message)
	}
	fmt.Println("Finalizando programa!")

}

func escrever(text string, chanel chan string) {

	for i := 0; i < 5; i++ {
		chanel <- text
		time.Sleep(time.Second)
	}

	close(chanel)
}
