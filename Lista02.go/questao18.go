package main

import f "fmt"

func main () {

	var v1, v2, v3 float64
	var cat, dia string

	f.Println("Digite o preço normal do DVD alugado: ")
	f.Scan(&v1)
	f.Println("Digite a categoria do DVD:")
	f.Scan(&cat)
	f.Println("Digite o dia da semana:")
	f.Scan(&dia)

	if dia == "segunda" || dia == "terça" || dia == "quinta" {
		v2 = v1*0.6
		if cat == "comum" {
			f.Printf("O preço pago é: %.2f ", v2)
		}else{
			v3 = 0.75*v1
			f.Printf("O preço pago é: %.2f", v3)
		}
	}else {
		v2 = v1
		if cat == "comum" {
			f.Printf("O preço pago é: %.2f", v2)
		}else {
			f.Printf("O preço pago é: %.2f ", v2*1.15)
		}
	}
	

}