package main

import f "fmt"

func main() {
	var condicao string
	var preco, pago float64

	f.Println("Insira o valor do produto:")
	f.Scan(&preco)

	f.Println("Selecione a condeção de pagamento:")
	f.Println("1- À vista, dinheiro ou cheque, 10% de desconto")
	f.Println("2-À vista, cartão de crédito, 5% de desconto")
	f.Println("3-Em 2 vezes, preço normal de etiqueta sem juros")
	f.Println("4-Em 3 vezes, preço normal de etiqueta + 10% de juros")
	f.Scan(&condicao)
	
	if preco <= 0 {
		return
	}

	if condicao == "1" {
       pago = preco*0.9
	   f.Printf("O preço pago será: %.2f", pago)
	}
	if condicao == "2" {
		pago = preco*0.95
		f.Printf("O preço pago será: %.2f", pago)
	}
	if condicao == "3" {
		pago = preco
		f.Printf("O preço pago será: %.2f", pago)
	}
	if condicao == "4" {
		pago = preco*1.1
		f.Printf("O preço pago será: %.2f", pago)
	}
}