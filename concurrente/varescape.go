package main

import (
	"fmt"
	"time"
)

//aplicando analisis de escape
func countdown(seconds *int) {

	for *seconds>0{
		time.Sleep(time.Second*1)
		*seconds -= 1
	}

}

func main() {
	count := 5
	go countdown(&count) //referencia en memoria
	for count > 0 {
		time.Sleep(time.Millisecond * 500)

		//leer valor de variable
		fmt.Println(count)
	}
}