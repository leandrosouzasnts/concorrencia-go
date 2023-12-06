package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Escrevendo 01")
	escrever("Escrevendo 02")
}

func escrever(text string) {

	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
