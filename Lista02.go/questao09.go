package main

import f "fmt"

func main () {

	var v, venda float64

	f.Println("Insira o valor da compra: ")
	f.Scan(&v)

	if  v <= 0 {
		f.Println("Valor de compra inválido")
	}
	if v < 10.0 {
		venda = (v*10.0)/3
		f.Printf("O valor da venda é: %.2f", venda)
	}
	if v >= 10.0 && v < 30.0 {
		venda = (v*2)
		f.Printf("O valor da venda é:%.2f", venda)
	}
	if v >= 30.0 && v < 50.0 {
		venda = (v*10)/6
		f.Printf("O valor da venda é: %.2f", venda)
	}
	if v >= 50.0 {
		venda = (v*10)/7
		f.Printf("O valor da venda é: %.2f", venda)
	}
}	