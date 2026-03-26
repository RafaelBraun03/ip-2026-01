package main

import f "fmt"

func main () {

	var x, y float64

	f.Println("Digite um número : ")
	f.Scan(&x)

	if x == 2.0 {
		f.Println("Número inválido")
	}else {
		y = 8/(2 - x)
		f.Println("f(x)=", y)
	}
}