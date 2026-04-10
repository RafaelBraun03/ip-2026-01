package main
import (
	"fmt"
)
type Pessoa struct {
	nome      string
	idade     int
	profissao string
	salario   int
}
func main() {
	var pes1 Pessoa
	// Especificacao da pes1
	pes1.nome = "Monica"
	pes1.idade = 35
	pes1.profissao = "Engenheiro de Software"
	pes1.salario = 7000
	// Imprimir as informacoes da pessoa1
	fmt.Println("nome: ", pes1.nome)
	fmt.Println("idade: ", pes1.idade)
	fmt.Println("profissao: ", pes1.profissao)
	fmt.Println("salario: ", pes1.salario)
}
