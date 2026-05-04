package main

import f "fmt"

func main() {
	var A [10]int
	contagem := make(map[int]int)
	for i:= 0; i <len(A); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&A[i])
		contagem[A[i]]++
	}
	encontrou :=false // devemos devinir essa variável para caso nenhum número repetir
	for n,qtd:= range contagem{
		if qtd >1 {
			f.Printf("O número %d repete %d vezes",n,qtd)
			f.Println()
			encontrou = true
		}
	}

	if !encontrou{
		f.Println("Não há número repetido")
	}
	
}