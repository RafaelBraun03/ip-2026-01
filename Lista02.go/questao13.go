package main

import f "fmt"

func main () {
	
	var n, a, b int

	f.Println("Digite um número de 3 digitos: ")
	f.Scan(&n)

	if n >= 100 && n < 1000 {
		a = n / 100
		b = (n - a*100)/10
		f.Println("O algarismo da dexena é: ", b)
	}else {
		f.Print("O número é inválido")
	}


}
