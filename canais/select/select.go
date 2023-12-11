package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "canal 2"
		}
	}()

	for {
		select { //Verifica qual canal tem Msg (retorno) para dar continuidade
		case <-canal1:
			fmt.Println(<-canal1)
		case <-canal2:
			fmt.Println(<-canal2)
		}
	}

	// for {
	// 	fmt.Println(<-canal1)
	// 	fmt.Println(<-canal2)
	// }
}
