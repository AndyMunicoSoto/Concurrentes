package main

import "fmt"

var turn int = 1

func P() {
	for {
		fmt.Println("P line1 NCS")
		fmt.Println("P line2 NCS")

		for turn != 1{
			//espereandi
		}

		fmt.Println("P line 1 CS")
		fmt.Println("P line 2 CS")
		 turn = 2
	}
}

func Q() {
	for {
		fmt.Println("Q line 1 CS")
		fmt.Println("Q line 2 CS")

		for turn != 2 {
			//esperando
		}
		turn = 1
	}
}

func main() {
	go P()
	Q()
}