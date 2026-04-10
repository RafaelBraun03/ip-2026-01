package main

import (
	"fmt"
)

// 1. Criação da struct
type Pessoa struct {
	Nome      string
	Altura    float64
	PesoIdeal float64
}

func main() {
	// 2. Criação do slice vazio de structs 'Pessoa'
	var pessoas []Pessoa

	fmt.Println("=== Calculadora de Peso Ideal ===")
	fmt.Println("Digite 'FIM' no nome para encerrar o programa.\n")

	// 3. Repetição para ler os dados
	for {
		var nome string
		var altura float64

		fmt.Print("Nome: ")
		fmt.Scan(&nome)

		// Condição de parada: se digitar FIM, quebra o laço
		if nome == "FIM" {
			break
		}

		fmt.Print("Altura (em metros, ex: 1.75): ")
		fmt.Scan(&altura)

		// 4. Cálculo do peso ideal
		pesoIdeal := (72.7 * altura) - 58.0

		// Preenchendo a struct e adicionando ao slice
		novaPessoa := Pessoa{
			Nome:      nome,
			Altura:    altura,
			PesoIdeal: pesoIdeal,
		}
		
		pessoas = append(pessoas, novaPessoa)
		fmt.Println("---")
	}

	// 5. Imprimindo o slice formatado
	fmt.Println("\n=== Lista de Pessoas Cadastradas ===")
	for _, p := range pessoas {
		// %.2f formata os números flutuantes para exibirem apenas 2 casas decimais
		fmt.Printf("Nome: %s | Altura: %.2fm | Peso Ideal: %.2fkg\n", p.Nome, p.Altura, p.PesoIdeal)
	}
}