package main

import f "fmt"

func main() {
	var A [10]int
	
	for i:= 0; i <len(A); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&A[i])	
	}
	menor := A[0]
	posicao:=0

	for i,v:= range A{
		if v < menor {
			menor = v
			posicao = i
		}
	}
	f.Printf("O menor elemento é %d e o seu índice é: %d", menor, posicao)
}