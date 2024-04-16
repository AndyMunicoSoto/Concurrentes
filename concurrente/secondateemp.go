package main

import (
	"fmt"
	"time"
)

var wanp bool = false
var wanq bool = false

func P() {
	for {
		fmt.Println("P line1 NCS")
		fmt.Println("P line2 NCS")

		for wanq != false{
			wanp = true
		}

		fmt.Println("P line1 CS")
		fmt.Println("P line2 CS")
		wanp = false

	}
}

func Q() {
	for {
		fmt.Println("Q line1 NCS")
		fmt.Println("Q line2 NCS")

		for wanp != false{
			wanq = true
		}

		fmt.Println("Q line1 CS")
		fmt.Println("Q line2 CS")
		wanq = false

	}
}


func main() {
	go P()
	go Q()
	time.Sleep(time.Millisecond*500)
}