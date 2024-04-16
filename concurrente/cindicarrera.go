package main

import (
	"fmt"
	"time"
)

//con referecnia porque comparten memoria
func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
	}
	fmt.Println("Finaliza stingy")
}

func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
	}
	fmt.Println("Finaliza spendy")
}

func main() {
	money := 100

	go stingy(&money)
	go spendy(&money)
	time.Sleep(time.Millisecond*100)
	fmt.Println("Saldo final de la cuenta: ", money)
}