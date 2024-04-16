package main

import "time"

//funcion propia
func menores(num int, R[] int) int{

	//obtener el nro de elementos del arrelo
	n :=  len(R)

	c := 0

	for i := 0; i<n; i++ {

		if num > R[i] {
			c++
		}
		
	}
	

	return c
}

func main() {
	C := [] int{5, 1, 3 ,2 , 7, 4, 9, 8, 10, 6}
	D := [] int{0,0,0,0,0,0,0,0,0,0}

	var myNumber, count int

	for i:=0; i<10;i++ {
		//funciones anonimas
		go func (j int)  {
			myNumber = C[j]
			count = menores(myNumber, C)
			D[count] = myNumber

		}(i)
	}

	time.Sleep(time.Second)

	//print D
	for _, v:=range D{
		print(v)
	}

	

}