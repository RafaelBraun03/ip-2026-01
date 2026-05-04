// Aqui é necessário fazer uma ordenaçao por inserção/ insertion sort
package main

import (
	"fmt"
	f "fmt"
)

func main() {
	var vetor [10]int

	f.Println("Insira 10 números inteiros")
	for i := 0; i < 10; i++ {
		var num int
		fmt.Scan(&num)

		pos := 0                           //Em qual índice (posição) esse novo num deve ser guardado?
		for pos < i && vetor[pos] < num {
			pos++
		}

		for j := i; j > pos; j-- {    //empurrar todo mundo uma casa para a direita.
			vetor[j] = vetor[j-1]  // ele começa o loop da última posição até chegar na posição em que deve ser aberto espaço
		}							// Lembrando: i corresponde a posição em quue há o primeiro espaço vazio(última posição)
		
		vetor[pos] = num
		fmt.Println(vetor) // apenas para visualizar
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", vetor[i])
	}
	fmt.Println()
}