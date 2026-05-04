package main

import (f "fmt"
		 "sort")

type Funcionario struct{
	numero int
	meses int
}

func main (){
	empregados := make([]Funcionario, 0, 100)
	var numero, meses int
	f.Println("Digite o Número do Empregado e os Meses de Trabalho.")
	f.Println("Para encerrar, digite 0 0.")

	for i:=0; i <100;i++{
		f.Printf("Funcionário %d ",i+1)
		f.Println()
		f.Println("Id:")
		f.Scan(&numero)
		f.Println("Tempo:")
		f.Scan(&meses)

		if numero==0 && meses==0 {
			break
		}
		novoempregado:= Funcionario{numero: numero,meses: meses}
		empregados = append(empregados, novoempregado)
	}

	//A função sort.Slice exige dois parâmetros:
    //O vetor que você quer ordenar (empregados).
    //A regra de ordenação (a função embutida).
	sort.Slice(empregados, func(i, j int) bool {
		return empregados[i].meses < empregados[j].meses 
	})
	
	// agora que o slice está ordenado em relação aos meses, basta pegar os 3 primeiros
	//caso seja digitado um número de funcionários menor que 3:
	limite:=3
	if len(empregados)<3{
		limite=len(empregados)
	}
	for i:=0;i < limite; i++{
		f.Printf("Funcionário %d   Tempo: %d meses", empregados[i].numero,empregados[i].meses)
		f.Println()
	}

}