package main

import f "fmt"

func main() {
	var h [10]float64

	for i:= 0; i <len(h); i++ {
		f.Printf("Digite a %dº altura:", i+1)
		f.Scan(&h[i])
	}
	soma :=0.0
	for _,v:= range h {
		soma += v
	}
	media := soma/10
	f.Println("Média = ",media)
	for j,x := range h {
		if x > media{
			f.Printf("Atleta %d --- Altura = %.2f",j+1,x)
			f.Println()
		}
	}
}