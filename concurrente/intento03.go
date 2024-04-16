package main

import "fmt"

var wantp bool = false
var wantq bool = false

//precondiciona es la entrada
func P() {
	for {
		//non critical
		fmt.Println("proceso P line NCS")
		wantp = true
		for wantq {
			//espera
		}
		fmt.Println("Proc. P line1 CS")
		wantp = false

	}
}

func Q() {
	for {
		//non critical
		fmt.Println("proceso Q line NCS")
		for wantp {
			//espera
		}
		fmt.Println("Proc. Q line1 CS")
		wantq = false
	}
}

func main() {
	go P()
	Q()
	
}
