package main

import f "fmt"

func main() {

	var Qtd int
	var media, soma, nota float32

	f.Println("Insira a quantidade de notas:")
	f.Scan(&Qtd)

	if Qtd <= 0 {
		f.Println("Quantidade inválida")
	}
	for i := 1; i <= Qtd; i++ {
		f.Printf("Digite a nota %d: ", i)
		f.Scan(&nota)
		soma += nota
	}
	media = soma / float32(Qtd)
	f.Println("A méidia é: ", media)
}
