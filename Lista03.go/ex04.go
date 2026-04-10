package main

import f "fmt"

func main() {
	var n int

	f.Println("Insira um número inteiro:")
	f.Scan(&n)

	if n <= 0 {
		f.Println("Número inválido")
	}
    quadrado := false
	for i := 1; i*i <= n; i++{
		if i*i == n {
			quadrado = true
			break
		}	
	}
	if quadrado {
		f.Printf("%d é um quadrado perfeito", n)
	}else {
		f.Printf("%d não é um quadrado perfeito", n)
	}


}