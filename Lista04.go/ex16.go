package main

import f "fmt"

func main() {
	var idade [50]int
	contagem := make(map[int]int)

	for i:= 0; i <len(idade); i++ {
		f.Printf("Digite o %dº valor:", i+1)
		f.Scan(&idade[i])
		contagem[idade[i]]++
	}
	maior:=0
	for _,qtd:= range contagem{
		if qtd > maior{
			maior = qtd
		}
	}
	for id,qtd:=range contagem{
		if qtd == maior{
			f.Printf("Moda=%d    Frequência=%d",id,qtd)
			f.Println()
		}
	}
}