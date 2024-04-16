package main

import "fmt"

var wantp bool = false
var wantq bool = false

func P() {
	for {
		//critical section
		fmt.Println("proceso P 1 NCS")
		wantp = true
		for wantq {
			wantp = false
			wantp = true
		}
		//cs
		fmt.Println("section critical P")
		wantp = false
	}
}

func Q() {
	for {
		//critical section
		fmt.Println("proceso Q NCS")
		wantq = true
		for wantp {
			wantq = false
			wantq = true
		}
		//cs
		fmt.Println("section critical Q")
	
		wantq = false
	}

}

func main() {
	go P()
	Q()
}
