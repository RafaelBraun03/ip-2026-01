package main

import (f "fmt"
	    s "strings")

func main () {

	var v1, v2 float64
	var resposta string

	const (v_ar =1750.00
	       v_pintura = 800.00
		   v_vidro = 1200.00
		   v_direcao = 2000.00)
		   
    f.Print("Digite o preço inicial de fábrica do carro: ")
	f.Scan(&v1)

	v2 = v1
	f.Println("Usando s para SIM e n para NAO, selecione as opções abaixo:")

	f.Printf("a) Ar condicionado (R$ %.2f)? (s/n): ", v_ar)
	f.Scan(&resposta)
	if s.ToLower(resposta) == "s" {
		v2 = v2 + v_ar
	}
	f.Printf("b) Pintura metálica (R$ %.2f)? (s/n): ", v_pintura)
	f.Scan(&resposta)
	if s.ToLower(resposta) =="s" {
		v2 = v2 + v_pintura
	}
	f.Printf("c) Vidro elétrico (R$ %.2f)? (s/n): ", v_vidro)
	f.Scan(&resposta)
	if s.ToLower(resposta) == "s" {
           v2 = v2 + v_vidro
	}
	f.Printf("d) Direção hidráulica (R$ %.2f)? (s/n): ", v_direcao)
	f.Scan(&resposta)
	if s.ToLower(resposta) == "s" {
		v2 = v2 + v_direcao
	}
	f.Printf("O valor final do carro é: %.2f", v2)

	}
		