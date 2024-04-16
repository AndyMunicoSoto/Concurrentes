package main

import (
	"fmt"
	"time"
)

var n int = 0
var flag bool = false

func P() {
	for flag == false {
		n = 1 - n
		fmt.Println("Proceso P valor de n = ",n," valor de flag = ", flag)
	}
}

func Q() {
	for flag == false {
		if n == 0 {
			flag = true
			fmt.Println("Proceso Q valor de n = ",n," valor de flag = ", flag)
		}
	}
}

func main() {
	go P()
	go Q()

	time.Sleep(time.Microsecond*100)
}