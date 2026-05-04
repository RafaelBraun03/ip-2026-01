package main 

import f "fmt"

func main() {
	var jogadas [20]int

	frequencia := make(map[int]int)

	for  i:= 0; i <len(jogadas); i++ {
		f.Printf("Digite o %dº valor :", i+1)
		f.Scan(&jogadas[i])
		if jogadas[i] >6 || jogadas[i]<1 {
			f.Println("Número inválido")
			return
		}
		frequencia[jogadas[i]]++
	}
	for n,qtd:= range frequencia{
		f.Printf("Número: %d     Frequência: %d", n, qtd)
		f.Println()
	}
}