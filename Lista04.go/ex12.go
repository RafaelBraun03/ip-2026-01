package main

import f "fmt"

func main() {
	var notas [15]float64
	var relativo float64
	contagem := make(map[float64]int)
	for i:= 0; i <len(notas); i++ {
		f.Printf("Digite a %dº nota:", i+1)
		f.Scan(&notas[i])
		contagem[notas[i]]++
	}
	
	for n,qtd := range contagem{
		relativo = 100*float64(qtd)/15
		f.Printf("NOTA:%.2f     FREQUÊNCIA ABSOLUTA:%d     FREQUÊNCIA RELATIVA:%.2f",n,qtd,relativo)
		f.Print("%")
		f.Println()
	}

}